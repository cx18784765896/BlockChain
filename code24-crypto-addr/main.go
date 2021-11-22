package main

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	//sha256
	hash256 := sha256.New()
	hash256.Write([]byte("Sirius"))
	byte256 := hash256.Sum(nil)
	fmt.Printf("SHA256:%x\n",byte256)
	//ripemd160
	hash160 := ripemd160.New()
	hash160.Write([]byte("Sirius"))
	byte160 := hash160.Sum(nil)
	fmt.Printf("Ripemd160:%x\n",byte160)
}
