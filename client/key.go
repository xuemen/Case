package main

import (
	"golang.org/x/crypto/openpgp"
	//"golang.org/x/crypto/openpgp/armor"
	"os"
	//"time"
)

var pubkey, seckey *openpgp.Entity

func getKeyByEmail(keyring openpgp.EntityList, email string) *openpgp.Entity {
	for _, entity := range keyring {
		for _, ident := range entity.Identities {
			if ident.UserId.Email == email {
				return entity
			}
		}
	}

	return nil
}

func ReadJSKey() bool {
	pubringFile, err := os.Open("key.pub")
	if err != nil {
		return false
	}

	defer pubringFile.Close()
	pubring, _ := openpgp.ReadArmoredKeyRing(pubringFile)
	pubkey = getKeyByEmail(pubring, conf.Email)

	secringFile, _ := os.Open("key.sec")
	defer secringFile.Close()
	secring, _ := openpgp.ReadArmoredKeyRing(secringFile)
	seckey = getKeyByEmail(secring, conf.Email)

	return true
}
