package cli

import (
	"flag"
	"fmt"
	"github.com/gookit/goutil/fsutil"
	"os"
)

var (
	AuthKeysFile string
)

func init() {
	flag.StringVar(&AuthKeysFile, "auth-keys-file", "", "Auth Keys json file path")
}

func Parse() {
	flag.Parse()
	validate()
}

func validate() {
	if !fsutil.IsFile(AuthKeysFile) {
		fmt.Println("Auth Keys json file not exist")
		os.Exit(1)
	}
}
