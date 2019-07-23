package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dappley/go-dappley/client"
	"github.com/dappley/go-dappley/common"
	"github.com/dappley/go-dappley/consensus"
	"github.com/dappley/go-dappley/core"
	"github.com/dappley/go-dappley/logic"
	"github.com/dappley/go-dappley/storage"
	logger "github.com/sirupsen/logrus"
)

const (
	genesisAddr           = "121yKAXeG4cw6uaGCBYjWk9yTWmMkhcoDD"
	genesisFilePath       = "conf/genesis.conf"
	defaultPassword       = "password"
	defaultTimeBetweenBlk = 5
	contractFunctionCall  = "{\"function\":\"record\",\"args\":[\"dEhFf5mWTSe67mbemZdK3WiJh8FcCayJqm\",\"4\"]}"
	contractFilePath      = "contract/test_contract.js"
)

var (
	password    = "testpassword"
	maxAccount  = 4
	currBalance = make(map[string]uint64)
	numOfTx     = 100
	numOfScTx   = 0
	time        = int64(1532392928)
)

type FileInfo struct {
	Height        int
	DifferentFrom int
	Db            *storage.LevelDB
}

type Key struct {
	Key     string `json:"key"`
	Address string `json:"address"`
}

type Keys struct {
	Keys []Key `json:"keys"`
}

type GeneralConfigs struct {
	NumOfNormalTx int
	NumOfScTx     int
}

func GenerateNewBlockChain(files []FileInfo, d *consensus.Dynasty, keys Keys, config GeneralConfigs) {
	bcs := make([]*core.Blockchain, len(files))
	addr := client.NewAddress(genesisAddr)
	numOfTx = config.NumOfNormalTx
	numOfScTx = config.NumOfScTx
	for i := range files {
		bc := core.CreateBlockchain(addr, files[i].Db, nil, 2000, nil, 1000000)
		bcs[i] = bc
	}

	for i, p := range d.GetProducers() {
		logger.WithFields(logger.Fields{
			"producer": p,
		}).Info("Producer:", i)
	}

	wm, err := logic.GetAccountManager(client.GetAccountFilePath())
	if err != nil {
		logger.Panic("Cannot get account manager.")
	}
	addrs := CreateAccount(wm)
	producer := client.NewAddress(d.ProducerAtATime(time))
	key := keys.getPrivateKeyByAddress(producer)
	logic.SetMinerKeyPair(key)

	//max, index := GetMaxHeightOfDifferentStart(files)
	//fund every miner
	parentBlks := make([]*core.Block, len(files))
	utxoIndexes := make([]*core.UTXOIndex, len(files))
	for i := range files {
		parentBlks[i], _ = bcs[i].GetTailBlock()
		utxoIndexes[i] = core.NewUTXOIndex(bcs[i].GetUtxoCache())
		for j := 0; j < len(d.GetProducers()); j++ {
			b := generateBlock(utxoIndexes[i], parentBlks[i], bcs[i], d, keys, []*core.Transaction{})
			bcs[i].AddBlockToDb(b)
			parentBlks[i] = b
		}
	}

	//fund from miner
	fundingBlock := generateFundingBlock(utxoIndexes[0], parentBlks[0], bcs[0], d, keys, addrs[0], key)
	for idx := range files {
		bcs[idx].AddBlockToDb(fundingBlock)
	}
	parentBlks[0] = fundingBlock

	//deploy smart contract
	scblock, scAddr := generateSmartContractDeploymentBlock(utxoIndexes[0], parentBlks[0], bcs[0], d, keys, addrs[0], wm)
	logger.Info("smart contract address:", scAddr.String())
	for idx := range files {
		bcs[idx].AddBlockToDb(scblock)
	}
	parentBlks[0] = scblock

	for i, file := range files {
		makeBlockChainToSize(utxoIndexes[i], parentBlks[i], bcs[i], file.Height, d, keys, addrs, wm, scAddr)
	}

}

func GetMaxHeightOfDifferentStart(files []FileInfo) (int, int) {
	max := 0
	index := 0
	for i, file := range files {
		if max < file.DifferentFrom {
			max = file.DifferentFrom
			index = i
		}
	}
	return max, index
}

func makeBlockChainToSize(utxoIndex *core.UTXOIndex, parentBlk *core.Block, bc *core.Blockchain, size int, d *consensus.Dynasty, keys Keys, addrs []client.Address, wm *client.AccountManager, scAddr client.Address) {

	tailBlk := parentBlk
	for tailBlk.GetHeight() < uint64(size) {
		txs := generateTransactions(utxoIndex, addrs, wm, scAddr)
		b := generateBlock(utxoIndex, tailBlk, bc, d, keys, txs)
		bc.AddBlockToDb(b)
		tailBlk = b
	}
	bc.GetDb().Put([]byte("tailBlockHash"), tailBlk.GetHash())
}

func generateBlock(utxoIndex *core.UTXOIndex, parentBlk *core.Block, bc *core.Blockchain, d *consensus.Dynasty, keys Keys, txs []*core.Transaction) *core.Block {
	producer := client.NewAddress(d.ProducerAtATime(time))
	key := keys.getPrivateKeyByAddress(producer)
	cbtx := core.NewCoinbaseTX(producer, "", parentBlk.GetHeight()+1, common.NewAmount(0))
	txs = append(txs, &cbtx)
	utxoIndex.UpdateUtxo(&cbtx)
	b := core.NewBlockWithTimestamp(txs, parentBlk, time, producer.String())
	hash := b.CalculateHashWithNonce(0)
	b.SetHash(hash)
	b.SetNonce(0)
	b.SignBlock(key, hash)
	time = time + defaultTimeBetweenBlk
	logger.WithFields(logger.Fields{
		"producer":  producer.String(),
		"timestamp": time,
		"blkHeight": b.GetHeight(),
	}).Info("Tool:Generating Block...")
	return b
}

func generateFundingBlock(utxoIndex *core.UTXOIndex, parentBlk *core.Block, bc *core.Blockchain, d *consensus.Dynasty, keys Keys, fundAddr client.Address, minerPrivKey string) *core.Block {
	logger.Info("generate funding Block")
	tx := generateFundingTransaction(utxoIndex, fundAddr, minerPrivKey)
	return generateBlock(utxoIndex, parentBlk, bc, d, keys, []*core.Transaction{tx})
}

func generateSmartContractDeploymentBlock(utxoIndex *core.UTXOIndex, parentBlk *core.Block, bc *core.Blockchain, d *consensus.Dynasty, keys Keys, fundAddr client.Address, wm *client.AccountManager) (*core.Block, client.Address) {
	logger.Info("generate smart contract deployment block")
	tx := generateSmartContractDeploymentTransaction(utxoIndex, fundAddr, wm)

	return generateBlock(utxoIndex, parentBlk, bc, d, keys, []*core.Transaction{tx}), tx.Vout[0].PubKeyHash.GenerateAddress()
}

func generateSmartContractDeploymentTransaction(utxoIndex *core.UTXOIndex, sender client.Address, wm *client.AccountManager) *core.Transaction {
	senderAccount := wm.GetAccountByAddress(sender)
	if senderAccount == nil || senderAccount.GetKeyPair() == nil {
		logger.Panic("Can not find sender account")
	}
	pubKeyHash, _ := client.NewUserPubKeyHash(senderAccount.GetKeyPair().PublicKey)

	data, err := ioutil.ReadFile(contractFilePath)
	if err != nil {
		logger.WithError(err).WithFields(logger.Fields{
			"file_path": contractFilePath,
		}).Panic("Unable to read smart contract file!")
	}
	contract := string(data)
	tx := newTransaction(sender, client.Address{}, senderAccount.GetKeyPair(), utxoIndex, pubKeyHash, common.NewAmount(1), common.NewAmount(10000), common.NewAmount(1), contract)
	utxoIndex.UpdateUtxo(tx)
	currBalance[sender.String()] -= 1
	return tx
}

func generateFundingTransaction(utxoIndex *core.UTXOIndex, fundAddr client.Address, minerPrivKey string) *core.Transaction {
	initFund := uint64(1000000)
	initFundAmount := common.NewAmount(initFund)
	minerKeyPair := client.GetKeyPairByString(minerPrivKey)
	pkh, _ := client.NewUserPubKeyHash(minerKeyPair.PublicKey)

	tx := newTransaction(minerKeyPair.GenerateAddress(), fundAddr, minerKeyPair, utxoIndex, pkh, initFundAmount, common.NewAmount(10000), common.NewAmount(1), "")
	utxoIndex.UpdateUtxo(tx)
	currBalance[fundAddr.String()] = initFund
	return tx
}

func generateTransactions(utxoIndex *core.UTXOIndex, addrs []client.Address, wm *client.AccountManager, scAddr client.Address) []*core.Transaction {
	pkhmap := getPubKeyHashes(addrs, wm)
	txs := []*core.Transaction{}
	for i := 0; i < numOfTx; i++ {
		contract := ""
		tx := generateTransaction(addrs, wm, utxoIndex, pkhmap, contract, scAddr)
		utxoIndex.UpdateUtxo(tx)
		txs = append(txs, tx)
	}
	for i := 0; i < numOfScTx; i++ {
		contract := contractFunctionCall
		tx := generateTransaction(addrs, wm, utxoIndex, pkhmap, contract, scAddr)
		utxoIndex.UpdateUtxo(tx)
		txs = append(txs, tx)
	}
	return txs
}

func getPubKeyHashes(addrs []client.Address, wm *client.AccountManager) map[client.Address]client.PubKeyHash {
	res := make(map[client.Address]client.PubKeyHash)
	for _, addr := range addrs {
		account := wm.GetAccountByAddress(addr)
		pubKeyHash, _ := client.NewUserPubKeyHash(account.GetKeyPair().PublicKey)
		res[addr] = pubKeyHash
	}
	return res
}

func generateTransaction(addrs []client.Address, wm *client.AccountManager, utxoIndex *core.UTXOIndex, pkhmap map[client.Address]client.PubKeyHash, contract string, scAddr client.Address) *core.Transaction {
	sender, receiver := getSenderAndReceiver(addrs)
	amount := common.NewAmount(1)
	senderAccount := wm.GetAccountByAddress(sender)
	if senderAccount == nil || senderAccount.GetKeyPair() == nil {
		logger.Panic("Can not find sender account")
	}
	if contract != "" {
		receiver = scAddr
	}
	tx := newTransaction(sender, receiver, senderAccount.GetKeyPair(), utxoIndex, pkhmap[sender], amount, common.NewAmount(10000), common.NewAmount(1), contract)
	currBalance[sender.String()] -= 1
	currBalance[receiver.String()] += 1

	return tx
}

func newTransaction(sender, receiver client.Address, senderKeyPair *client.KeyPair, utxoIndex *core.UTXOIndex, senderPkh client.PubKeyHash, amount *common.Amount, gasLimit *common.Amount, gasPrice *common.Amount, contract string) *core.Transaction {
	utxos, _ := utxoIndex.GetUTXOsByAmount([]byte(senderPkh), amount)

	sendTxParam := core.NewSendTxParam(sender, senderKeyPair, receiver, amount, common.NewAmount(0), gasLimit, gasPrice, contract)
	tx, err := core.NewUTXOTransaction(utxos, sendTxParam)

	if err != nil {
		logger.WithError(err).Panic("Create transaction failed!")
	}

	return &tx
}

func getSenderAndReceiver(addrs []client.Address) (sender, receiver client.Address) {
	for i, addr := range addrs {
		if currBalance[addr.String()] > 1000 {
			sender = addr
			if i == maxAccount {
				receiver = addrs[0]
			} else {
				receiver = addrs[i+1]
			}
			return
		}
	}
	for key, val := range currBalance {
		logger.WithFields(logger.Fields{
			"addr": key,
			"val":  val,
		}).Info("Current Balance")
	}
	logger.Panic("getSenderAndReceiver failed")
	return
}

func CreateRandomTransactions([]client.Address) []*core.Transaction {
	return nil
}

func CreateAccount(wm *client.AccountManager) []client.Address {

	addresses := wm.GetAddresses()
	numOfAccounts := len(addresses)
	for i := numOfAccounts; i < maxAccount; i++ {
		_, err := logic.CreateAccountWithpassphrase(password)
		if err != nil {
			logger.WithError(err).Panic("Cannot create new account.")
		}
	}

	addresses = wm.GetAddresses()
	logger.WithFields(logger.Fields{
		"addresses": addresses,
	}).Info("Accounts are created")
	return addresses
}

func LoadPrivateKey() Keys {
	jsonFile, err := os.Open("conf/key.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var keys Keys

	json.Unmarshal(byteValue, &keys)

	return keys
}

func (k Keys) getPrivateKeyByAddress(address client.Address) string {
	for _, key := range k.Keys {
		if key.Address == address.String() {
			return key.Key
		}
	}
	return ""
}
