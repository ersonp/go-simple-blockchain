package blockchain

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	mrand "math/rand"
	"strings"
	"time"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	host "github.com/libp2p/go-libp2p-core/host"
	net "github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	pstore "github.com/libp2p/go-libp2p-core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
)

// Address is used to store the address info of the host
type Address struct {
	HostAddr ma.Multiaddr
	HostID   string
	Addrs    []ma.Multiaddr
	FullAddr ma.Multiaddr
}

// Metrics is used to store
type Metrics struct {
	PeerID    peer.ID
	Addrs     []ma.Multiaddr
	Protocals []string
	PublicKey crypto.PubKey
	Ping      time.Duration
}

// HostAddressField is used to share the Adress data of the host
// throughtout the application
var HostAddressField Address

// PeerMetrics is used to share data of all the peers of the host
// throughtout the application
var PeerMetrics []Metrics

// MakeBasicHost creates a LibP2P host with a random peer ID listening on the
// given multiaddress. It will use secio if secio is true.
func MakeBasicHost(listenPort int, secio bool, randseed int64) (host.Host, error) {

	var a Address
	// If the seed is zero, use real cryptographic randomness. Otherwise, use a
	// deterministic randomness source to make generated keys stay the same
	// across multiple runs
	var r io.Reader
	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	// Generate a key pair for this host. We will use it
	// to obtain a valid host ID.
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		libp2p.Identity(priv),
	}

	basicHost, err := libp2p.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	// Build host multiaddress
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", basicHost.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	addrs := basicHost.Addrs()
	var addr ma.Multiaddr
	// select the address starting with "ip4"
	for _, i := range addrs {
		if strings.HasPrefix(i.String(), "/ip4") {
			addr = i
			break
		}
	}
	fullAddr := addr.Encapsulate(hostAddr)
	log.Printf("I am %s\n", fullAddr)
	if secio {
		log.Printf("Now run \"go run ./ -l %d -d %s -secio\" on a different terminal\n", listenPort+1, fullAddr)
	} else {
		log.Printf("Now run \"go run ./ -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
	}

	a.HostAddr = hostAddr
	a.HostID = basicHost.ID().Pretty()
	a.Addrs = basicHost.Addrs()
	a.FullAddr = fullAddr
	HostAddressField = a

	return basicHost, nil
}

// HandleStream is used to handle stream
func HandleStream(s net.Stream) {

	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	UpdateBlockchain(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
}

// UpdateBlockchain is used to keep the local blockchain updated with the newtork
func UpdateBlockchain(rw *bufio.ReadWriter) {
	// Create a thread to read and write data.
	go WriteBlockchain(rw)
	go ReadBlockchain(rw)
}

// ReadBlockchain is use to read the blockchain on the network
func ReadBlockchain(rw *bufio.ReadWriter) {

	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if str == "" {
			return
		}
		if str != "\n" {

			chain := make([]Block, 0)
			if err := json.Unmarshal([]byte(str), &chain); err != nil {
				log.Fatal(err)
			}

			mutex.Lock()
			if len(chain) > len(Blockchain) {
				Blockchain = chain
				bytes, err := json.MarshalIndent(Blockchain, "", "  ")
				if err != nil {

					log.Fatal(err)
				}
				// Green console color: 	\x1b[32m
				// Reset console color: 	\x1b[0m
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			mutex.Unlock()
		}
	}
}

// WriteBlockchain is use to write a block to the blockchain after validation
func WriteBlockchain(rw *bufio.ReadWriter) {

	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		bytes, err := json.Marshal(Blockchain)
		if err != nil {
			log.Println(err)
		}
		mutex.Unlock()

		mutex.Lock()
		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		rw.Flush()
		mutex.Unlock()

	}

}

// ConnectPeer is used to extablish a connection with a peer
func ConnectPeer(ha host.Host, target string) {
	// The following code extracts target's peer ID from the
	// given multiaddress
	ipfsaddr, err := ma.NewMultiaddr(target)
	if err != nil {
		log.Fatalln(err)
	}
	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		log.Fatalln(err)
	}

	peerid, err := peer.IDB58Decode(pid)
	if err != nil {
		log.Fatalln(err)
	}

	// Decapsulate the /ipfs/<peerID> part from the target
	// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := ma.NewMultiaddr(
		fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	ha.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

	a := ha.Peerstore().Peers()

	for _, i := range a {
		var m Metrics
		m.PeerID = i
		m.Addrs = ha.Peerstore().Addrs(i)
		m.Protocals, _ = ha.Peerstore().SupportsProtocols(i)
		m.PublicKey = ha.Peerstore().PubKey(i)
		m.Ping = ha.Peerstore().LatencyEWMA(i)
		PeerMetrics = append(PeerMetrics, m)
	}

	log.Println("opening stream")
	// make a new stream from host B to host A
	// it should be handled on host A by the handler we set above because
	// we use the same /p2p/1.0.0 protocol
	s, err := ha.NewStream(context.Background(), peerid, "/p2p/1.0.0")
	if err != nil {
		log.Fatalln(err)
	}
	// Create a buffered stream so that read and writes are non blocking.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	UpdateBlockchain(rw)

}
