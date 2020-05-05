// zlmkeygen is a minimal key pair generation tool for ZLM2.
package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
)

// Pair defines a Ed25519 key pair.
type Pair struct {
	PublicKey []byte
	SecretKey []byte
}

// Generate generates a new Ed25519 key pair.
func Generate() (*Pair, error) {
	var kp Pair
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	kp.PublicKey = pubKey[:]
	kp.SecretKey = privKey[:]
	return &kp, nil
}

// Marshal the Ed25519 key pair as a JSON string.
func (kp *Pair) Marshal() (string, error) {
	buf, err := json.MarshalIndent(kp, "", "  ")
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: error: %s\n", os.Args[0], err)
	os.Exit(1)
}

func main() {
	keyPair, err := Generate()
	if err != nil {
		fatal(err)
	}
	jsn, err := keyPair.Marshal()
	if err != nil {
		fatal(err)
	}
	fmt.Println(jsn)
}
