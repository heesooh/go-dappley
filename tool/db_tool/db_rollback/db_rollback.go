package main

import (
	"fmt"
    "errors"
    "os"
    "flag"
    "encoding/hex"
    "github.com/dappley/go-dappley/storage"
    "github.com/dappley/go-dappley/core/block"
    "github.com/dappley/go-dappley/core/transaction"
    "github.com/dappley/go-dappley/util"
    "github.com/dappley/go-dappley/common/hash"
    "encoding/json"
    "github.com/dappley/go-dappley/core/scState"
    "github.com/dappley/go-dappley/core/utxo"
    "github.com/dappley/go-dappley/logic/lutxo"
    "github.com/dappley/go-dappley/logic/transactionpool"
    logger "github.com/sirupsen/logrus"
    "google.golang.org/grpc/status"
)

var tipKey = []byte("tailBlockHash")

var (
    ErrBlockDoesNotExist       = errors.New("block does not exist in db")
    ErrTailHashDoesNotExist    = errors.New("tail hash does not exist in db")
    ErrTargetHeightNotValid    = errors.New("target height is not valid")
    ErrFileNotExist            = errors.New("File does not exist")
)

//command names
const (
    getBlockByHeight   = "getBlockByHeight"
    getTailBlock       = "getTailBlock"
    rollback           = "rollback"
    help               = "help"
)

//flag names
const (
    flagHeight         = "height"
    flagDatabase       = "file"
)

//command list
var cmdList = []string{
    getBlockByHeight,
    getTailBlock,
    rollback,
    help,
}

type valueType int

//type enum
const (
    valueTypeString = iota
    valueTypeUint64
)

type flagPars struct {
    name         string
    defaultValue interface{}
    valueType    valueType
    usage        string
}

//description of each function
var descrip = map[string]string{
    getBlockByHeight : "print the block information of the target height",
    getTailBlock     : "print the tail block information",
    rollback         : "rollback the blockchain in database to the target height",
}

//configure input parameters/flags for each command
var cmdFlagsMap = map[string][]flagPars{
    getBlockByHeight: {
        flagPars{
            flagDatabase,
            "default.db",
            valueTypeString,
            "database name. Eg. default.db",
        },
        flagPars{
            flagHeight,
            uint64(0),
            valueTypeUint64,
            "height. Eg. 1000",
        },
    },
    getTailBlock: {
        flagPars{
            flagDatabase,
            "default.db",
            valueTypeString,
            "database name. Eg. default.db",
        },
    },
    rollback: {
        flagPars{
            flagDatabase,
            "default.db",
            valueTypeString,
            "database name. Eg. default.db",
        },
        flagPars{
            flagHeight,
            uint64(0),
            valueTypeUint64,
            "height. Eg. 1000",
        },
    },
}

type commandHandler func(flags cmdFlags)

//map the callback function to each command
var cmdHandlers = map[string]commandHandler{
    getBlockByHeight:  getBlockByHeightCmdHandler,
    getTailBlock:      getTailBlockCmdHandler,
    rollback:          rollbackCmdHandler,
    help:              helpCmdHandler,
}

//map key: flag name   map defaultValue: flag defaultValue
type cmdFlags map[string]interface{}

func main() {
    args := os.Args[1:]

    if len(args) < 1 {
        printUsage()
        return
    }
    cmdFlagSetList := map[string]*flag.FlagSet{}
    //set up flagset for each command
    for _, cmd := range cmdList {
        fs := flag.NewFlagSet(cmd, flag.ContinueOnError)
        cmdFlagSetList[cmd] = fs
    }
    cmdFlagValues := map[string]cmdFlags{}
    //set up flags for each command
    for cmd, pars := range cmdFlagsMap {
        cmdFlagValues[cmd] = cmdFlags{}
        for _, par := range pars {
            switch par.valueType {
            case valueTypeString:
                cmdFlagValues[cmd][par.name] = cmdFlagSetList[cmd].String(par.name, par.defaultValue.(string), par.usage)
            case valueTypeUint64:
                cmdFlagValues[cmd][par.name] = cmdFlagSetList[cmd].Uint64(par.name, par.defaultValue.(uint64), par.usage)
            }
        }
    }
    cmdName := args[0]

    cmd := cmdFlagSetList[cmdName]
    if cmd == nil {
        fmt.Println("\nError:", cmdName, "is an invalid command")
        printUsage()
    } else {
        err := cmd.Parse(args[1:])
        if err != nil {
            return
        }
        if cmd.Parsed() {
            cmdHandlers[cmdName](cmdFlagValues[cmdName])
        }
    }
}

func printUsage() {
    fmt.Println("Usage:")
    for _, cmd := range cmdList {
        fmt.Println(" ", cmd)
    }
    fmt.Println("Note: Use the command 'help' to get the command usage in details")
}

func getBlockByHeightCmdHandler(flags cmdFlags) {
    dbname := *(flags[flagDatabase].(*string))
    height := *(flags[flagHeight].(*uint64))

    db, err := LoadDBFile(dbname)
    if err != nil {
        fmt.Println("Error: File does not exist!")
        return
    }
    defer db.Close()

    block, err := GetBlockByHeight(db, height)
    if err != nil {
        fmt.Println("Error: ",status.Convert(err).Message())
        return
    }
    fmt.Printf("Current database name is %s\n", dbname)
    PrintBlock(block)
}

func getTailBlockCmdHandler(flags cmdFlags) {
    dbname := *(flags[flagDatabase].(*string))
    db, err := LoadDBFile(dbname)
    if err != nil {
        fmt.Println("Error: File does not exist!")
        return
    }
    defer db.Close()

    block, err := GetTailBlock(db)
    if err != nil {
        fmt.Println("Error: ",status.Convert(err).Message())
    }
    fmt.Printf("Current database name is %s\n", dbname)
    fmt.Println("tail block info: ")
    PrintBlock(block)
}

func rollbackCmdHandler(flags cmdFlags) {
    dbname := *(flags[flagDatabase].(*string))
    targetHeight := *(flags[flagHeight].(*uint64))
    db, err := LoadDBFile(dbname)
    if err != nil {
        fmt.Println("Error: File does not exist!")
        return
    }
    defer db.Close()

    err = RollBack(db, targetHeight)
    if err != nil {
        fmt.Println("Error: ",status.Convert(err).Message())
        return
    }
    fmt.Printf("Current database name is %s\n", dbname)
    fmt.Println("Rollback is done!")
    tailblk, err := GetTailBlock(db)
    if err != nil {
        fmt.Println("Error: ",status.Convert(err).Message())
        return
    }
    fmt.Println("Current tail block height is", tailblk.GetHeight())
}

func helpCmdHandler(flag cmdFlags) {
    for cmd, pars := range cmdFlagsMap {
        fmt.Println("\n-----------------------------------------------------------------")
        fmt.Printf("Command: %s\n", cmd)
        fmt.Printf("Description: %s\n", descrip[cmd])
        fmt.Printf("Usage Example: ./db_rollback %s", cmd)
        for _, par := range pars {
            fmt.Printf(" -%s", par.name)
            if par.name == flagDatabase {
                fmt.Printf(" default.db ")
                continue
            }
            if par.name == flagHeight {
                fmt.Printf(" 10 ")
                continue
            }
        }
    }
    fmt.Println()
}

//----------------------------------- help function ----------------------------------------//
func isDbExist(filename string) bool {
    _, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func LoadDBFile(filename string) (storage.Storage, error) {
    isExist := isDbExist(filename)
    if !isExist {
        return nil, ErrFileNotExist
    }
    db := storage.OpenDatabase(filename)
    return db, nil
}

func GetBlockByHeight(db storage.Storage, height uint64) (*block.Block, error) {
    h, err := db.Get(util.UintToHex(height))
    if err != nil {
        return nil, ErrBlockDoesNotExist
    }
    rawBytes, err := db.Get(h)

    return block.Deserialize(rawBytes), nil
}

func GetTailBlock(db storage.Storage) (*block.Block, error) {
    hash, err := db.Get(tipKey)
    if err != nil {
        return nil, ErrTailHashDoesNotExist
    }
    rawBytes, err := db.Get(hash)

    return block.Deserialize(rawBytes), nil
}


func RevertUtxoAndScState(db storage.Storage, targetHeight uint64) (*lutxo.UTXOIndex, *scState.ScState, error) {
    utxocache := utxo.NewUTXOCache(db)
    index := lutxo.NewUTXOIndex(utxocache)
    scState := scState.LoadScStateFromDatabase(db)
    tailBlock, err := GetTailBlock(db)
    if err != nil {
        return nil, nil, err
    }
    tailHeight := tailBlock.GetHeight()

    if (targetHeight < 0 || targetHeight > tailHeight) {
        return nil, nil, ErrTargetHeightNotValid
    }
    if tailHeight == targetHeight {
        return index, scState, nil
    }

    for h := tailHeight; h > targetHeight; h-- {
        block, err := GetBlockByHeight(db, h)
        if err != nil {
            return nil, nil, err
        }
        //revert utxo
        err = index.UndoTxsInBlock(block, db)
        if err != nil {
            logger.WithError(err).WithFields(logger.Fields{
                "hash": block.GetHash(),
            }).Warn("failed to calculate previous state of UTXO index for the block")
            return nil, nil, err
        }
        //revert scstate
        err = scState.RevertState(db, block.GetHash())
        if err != nil {
            logger.WithError(err).WithFields(logger.Fields{
                "hash": block.GetHash(),
            }).Warn("failed to calculate previous state of scState for the block")
            return nil, nil, err
        }
    }

    return index, scState, nil
}

//rollback db to certain height 
func RollBack(db storage.Storage, targetHeight uint64) error {
    utxo, scState, err := RevertUtxoAndScState(db, targetHeight)
    if err != nil {
        logger.Error("fail to revert utxo and scState")
        return err
    }
    rollBackUtxo := utxo.DeepCopy()
    rollScState := scState.DeepCopy()
    txPool := transactionpool.LoadTxPoolFromDatabase(db, nil, 128000)
    tailBlock, err := GetTailBlock(db)
    if err != nil {
        logger.Error("fail to get tail block")
        return err
    }
    tailHeight := tailBlock.GetHeight()
    for h := tailHeight; h > targetHeight; h-- {
        block, _ := GetBlockByHeight(db, h)
        for _, tx := range block.GetTransactions() {
            adaptedTx := transaction.NewTxAdapter(tx)
            if !adaptedTx.IsCoinbase() && !adaptedTx.IsRewardTx() && !adaptedTx.IsGasRewardTx() && !adaptedTx.IsGasChangeTx() {
                txPool.Rollback(*tx)
            }
        }
        err = db.Del(block.GetHash())
        if err != nil {
            logger.Error("fail to delete hash-block pairs")
            return err
        }
        err = db.Del(util.UintToHex(h))
        if err != nil {
            logger.Error("fail to delete height-hash pairs")
            return err
        }
    }

    db.EnableBatch()
    defer db.DisableBatch()

    tailHash, _ := db.Get(util.UintToHex(targetHeight))
    err = db.Put(tipKey, tailHash)
    if err != nil {
        logger.Error("fail to set tail hash")
        return err
    }

    txPool.SaveToDatabase(db)
    rollBackUtxo.Save()
    rollScState.SaveToDatabase(db)
    db.Flush()

    return nil
}

func PrintBlock(b *block.Block) {
    
    encodedBlock := map[string]interface{}{
        "Header": map[string]interface{}{
            "Hash":      b.GetHash().String(),
            "Prevhash":  b.GetPrevHash().String(),
            "Timestamp": b.GetTimestamp(),
            "Producer":  b.GetProducer(),
            "height":    b.GetHeight(),
        },
        "Transactions": tx_pretty_string(b),
    }

    blockinfo, err := json.MarshalIndent(encodedBlock, "", "  ")
    if err != nil {
        fmt.Println("Error:", err.Error())
    }

    fmt.Println(string(blockinfo))
    fmt.Println("\n")
}

func tx_pretty_string(b *block.Block) []map[string]interface{} {
    var encodedTransactions []map[string]interface{}

        for _, transaction := range b.GetTransactions() {

            var encodedVin []map[string]interface{}
            for _, vin := range transaction.Vin {
                encodedVin = append(encodedVin, map[string]interface{}{
                    "Vout":      vin.Vout,
                    "Signature": hex.EncodeToString(vin.Signature),
                    "PubKey":    string(vin.PubKey),
                })
            }

            var encodedVout []map[string]interface{}
            for _, vout := range transaction.Vout {
                encodedVout = append(encodedVout, map[string]interface{}{
                    "Value":      vout.Value,
                    "PubKeyHash": hex.EncodeToString(vout.PubKeyHash),
                    "Contract":   vout.Contract,
                })
            }

            encodedTransaction := map[string]interface{}{
                "ID":   hash.Hash(transaction.ID).String(),
                "Vin":  encodedVin,
                "Vout": encodedVout,
            }
            encodedTransactions = append(encodedTransactions, encodedTransaction)
        }

        return encodedTransactions
}



