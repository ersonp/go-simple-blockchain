package main

import (
	"flag"
	"log"
	"time"

	bc "github.com/ersonp/go-simple-blockchain/blockchain"
	golog "github.com/ipfs/go-log"
	gologging "github.com/whyrusleeping/go-logging"
)

func main() {
	t := time.Now()
	genesisBlock := bc.Block{}
	genesisBlock = bc.Block{
		Index:     0,
		Timestamp: t.String(),
		BPM:       0,
		Hash:      bc.CalculateHash(genesisBlock),
		PrevHash:  "",
	}

	bc.Blockchain = append(bc.Blockchain, genesisBlock)

	// LibP2P code uses golog to log messages. They log with different
	// string IDs (i.e. "swarm"). We can control the verbosity level for
	// all loggers with:
	golog.SetAllLoggers(gologging.INFO) // Change to DEBUG for extra info

	// Parse options from the command line
	listenF := flag.Int("l", 40000, "wait for incoming connections")
	target := flag.String("d", "", "target peer to dial")
	secio := flag.Bool("secio", false, "enable secio")
	seed := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	// Make a host that listens on the given multiaddress
	ha, err := bc.MakeBasicHost(*listenF, *secio, *seed)
	if err != nil {
		log.Fatal(err)
	}

	if *target == "" {
		log.Println("listening for connections")
		// Set a stream handler on host A. /p2p/1.0.0 is
		// a user-defined protocol name.
		ha.SetStreamHandler("/p2p/1.0.0", bc.HandleStream)

	} else {
		ha.SetStreamHandler("/p2p/1.0.0", bc.HandleStream)

		bc.ConnectPeer(ha, *target)

	}

	httpPort := *listenF + 1000
	// run the gorilla/mux server
	run(httpPort)
	// select {}
}
