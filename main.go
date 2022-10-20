package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/codahale/sss"
	"os"
)

func main() {

	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		fmt.Printf("GenerateSeed: unexpected error: %v", err)
	}
	extKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Printf("NewMaster: unexpected error: %v", err)
	}

	fmt.Println(extKey.String())
	pubKeyPoc, err := extKey.Neuter()
	if err != nil {
		fmt.Printf("Neuter: unexpected error: %v", err)
	}

	fmt.Println(pubKeyPoc)

	message:=  extKey.String()


	n1:=3
	k1:=3

	n := byte(n1)
	k := byte(k1)

	fmt.Printf("Message:\t%s\n\n",message)
	fmt.Printf("Policy. Any %d from %d\n\n",k1,n1)

	if (k1>n1) {
		fmt.Printf("Cannot do this, as k greater than n")
		os.Exit(0)
	}

	shares, _:= sss.Split(n, k, []byte(message))

	fmt.Println("shares ", shares)
	subset := make(map[byte][]byte, k)
	for x, y := range shares {
		fmt.Printf("Share:\t%d\t%s\n",x,hex.EncodeToString(y))
		subset[x] = y
		if len(subset) == int(k) {
			break
		}
	}

	reconstructed := string(sss.Combine(subset))
	fmt.Printf("\nReconstructed: %s\n",reconstructed)

}
