package main

import (
	"crypto/rand"
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"sync"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func ok(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type BitcoinPair struct {
	*btcutil.WIF
}

func NewBitcoinPair(key []byte) *BitcoinPair {
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), key)
	wif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
	ok(err)
	return &BitcoinPair{WIF: wif}

}

func (b *BitcoinPair) GetPubKey() string {
	pubKeyStr, err := btcutil.NewAddressPubKey(b.SerializePubKey(), &chaincfg.MainNetParams)
	ok(err)
	return pubKeyStr.EncodeAddress()
}

func (b *BitcoinPair) GetPrivKey() string {
	return b.String()
}

func main() {
	help := flag.Bool("help", false, "")
	format := flag.String("format", "default", "")
	random := flag.Bool("random", false, "")
	flag.Bool("seq", true, "")
	flag.Parse()
	args := flag.Args()

	var err error
	var startID, count int64
	var prefix string
	count = 1000

	n := len(args)

	if *random {
		if *format == "default" {
			*format = "PRIVKEY PUBKEY"
		}
		if n >= 1 {
			count, err = strconv.ParseInt(args[0], 10, 64)
			ok(err)
		}
	} else {
		if *format == "default" {
			*format = "ID PUBKEY"
		}

		if n == 0 {
			fmt.Println(fmt.Sprintf(`Usage:
  addrgen [-sequence] [-format "ID,PUBKEY"] PREFIX [START_ID=1] [COUNT=1000]
  addrgen -random [-format "PRIVKEY,PUBKEY"] [COUNT=1000]

Options:
  -sequence       the private key is SHA256($PREFIX$ID)
  -random         generate in random
  -format         available fields ID PRIVKEY PUBKEY BASE
  -help           show help`))
			if *help {
				os.Exit(1)
			}
			return
		}
		if n > 0 {
			prefix = args[0]
			startID = 1
		}
		if n > 1 {
			startID, err = strconv.ParseInt(args[1], 10, 64)
			ok(err)
		}
		if n > 2 {
			count, err = strconv.ParseInt(args[2], 10, 64)
			ok(err)
		}
	}

	pattern := regexp.MustCompile("(PUBKEY|PRIVKEY|ID|BASE)")
	ch := make(chan int64)
	go func() {
		for id := startID; id < startID+count; id++ {
			ch <- id
		}
		close(ch)
	}()

	m := runtime.GOMAXPROCS(0)
	m64 := int64(m)

	pairs := make([]chan string, m)
	for i := 0; i < m; i++ {
		pairs[i] = make(chan string, 2)
	}

	go func() {
		wg := sync.WaitGroup{}
		wg.Add(m)
		for t := 0; t < m; t++ {
			go func() {
				defer wg.Done()
				for id := range ch {
					var buf []byte
					if *random {
						buf = make([]byte, 32)
						rand.Read(buf)
					} else {
						buf = []byte(fmt.Sprintf("%s%d", prefix, id))
					}
					key := sha256.Sum256(buf)
					pair := NewBitcoinPair(key[:])

					r := pattern.ReplaceAllStringFunc(*format, func(m string) string {
						switch m {
						case "PUBKEY":
							return pair.GetPubKey()
						case "PRIVKEY":
							return pair.GetPrivKey()
						case "ID":
							return strconv.FormatInt(id, 10)
						case "BASE":
							return string(buf)
						default:
							return m
						}
					})

					pairs[id%m64] <- r
				}
			}()
		}
		wg.Wait()
		return
	}()

	for i := startID; i < startID+count; i++ {
		r := <-pairs[i%m64]
		fmt.Println(r)
	}
}
