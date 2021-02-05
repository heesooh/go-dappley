package main

import (
	"fmt"
	"os"
	"errors"
	"flag"
    "strconv"
    "io/ioutil"
    "strings"
    "github.com/dappley/go-dappley/storage"
    networkpb "github.com/dappley/go-dappley/network/pb"
    "github.com/golang/protobuf/proto"
    logger "github.com/sirupsen/logrus"
    "github.com/dappley/go-dappley/network/networkmodel"
)

const (
	syncPeersKey = "SyncPeers"
)

var (
	ErrFileNotExist = errors.New("File does not exist")
)

func main() {
	args := os.Args[1:]
	if (len(args) < 1 || (len(args) >= 1 && args[0] != "-file")) {
		printUsage()
		return
	}
	var filePath string
	flag.StringVar(&filePath, "file", "default.db", "default db file path")
	flag.Parse()

	db, err := LoadDBFile(filePath)
	if err != nil {
		logger.WithError(err).Warn("Failed to load the file")
		return
	}
	defer db.Close()
	logger.Infof("Current database name is %s", filePath)

	peersBytes ,err := db.Get([]byte(syncPeersKey))

	if err != nil {
		logger.WithError(err).Warn("syncPeersKey doesn't exist in the database!")
		return
	}
	err = LoadPeerInfo2Config(peersBytes)
	if err != nil {
		return
	}
	err = db.Del([]byte(syncPeersKey))
	if err != nil {
		logger.WithError(err).Warn("Failed to delete the key-value pair!")
	}
}

func printUsage() {
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("Usage: split the database and load the peer info into peerinfo_config.conf")
	fmt.Println("Usage example: ./db_split -file default.db")
}

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

func addr_content(p networkmodel.PeerInfo) string {
	var addstr []string

	for _, addr := range p.Addrs {
		addstr = append(addstr, "\"" + addr.String() + "\"")
	}
	return "[" + strings.Join(addstr, ",") + "]"
}

func LoadPeerInfo2Config(peersBytes []byte) error {
	peerListPb := &networkpb.ReturnPeerList{}

	if err := proto.Unmarshal(peersBytes, peerListPb); err != nil {
		logger.WithError(err).Warn("parse Peerlist failed.")
		return err
	}

	var config_content string = "peers_config{\n"
	for index, peerPb := range peerListPb.GetPeerList() {
		peerInfo := networkmodel.PeerInfo{}
		if err := peerInfo.FromProto(peerPb); err != nil {
			logger.WithError(err).Warn("parse PeerInfo failed.")
			return err
		}

		logger.WithFields(logger.Fields{
			"peer_id": peerInfo.PeerId,
			"addr":    addr_content(peerInfo),
			"latency": peerInfo.Latency,
		}).Info("Loading syncPeers from database...")

		peername := "\tpeer_" + strconv.Itoa(index) + "{\n"
		pid := "\t\tpeer_id: \"" + peerInfo.PeerId.String() + "\"\n"
		paddr := "\t\taddr: \"" + addr_content(peerInfo) + "\"\n"
		platency := ""
		if peerInfo.Latency != nil {
			platency = ("\t\tlatency: " + fmt.Sprintf("%f", *peerInfo.Latency) + "\n")
		}
		config_content += (peername + pid + paddr + platency + "\t}\n")
	}
	config_content += "}"

	filename := "peerinfo_config.conf"
	ioutil.WriteFile(filename, []byte(config_content), 0644)
	return nil
}

