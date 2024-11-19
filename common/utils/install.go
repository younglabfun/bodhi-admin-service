package utils

import (
	"fmt"
	"os"
)

const LOCKFILE = "etc/install.lock"

func CheckLock() bool {
	_, err := os.Stat(LOCKFILE)
	if err == nil {
		// lock exist, installed
		return true
	}
	return false
}

func CreatLock() error {
	lockTime := GetDateTime()
	err := os.WriteFile(LOCKFILE, []byte(lockTime), 0666)
	if err != nil {
		return err
	}
	return nil
}

var (
	reset string = "n"
)

// CheckConf if conf file doesn't exist, create it
func CheckConf(file string) (string, error) {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		err = CreateFileInPath(file)
		reset = "y"
		if err != nil {
			return reset, err
		}
	} else {
		fmt.Printf("conf file exist, reset it? default no (y/n)")
		reset = ScanInput("n")

		if reset == "y" {
			err = CreateFile(file)
			if err != nil {
				return reset, err
			}
		}
	}
	return reset, nil
}

func ScanInput(def string) string {
	var input string
	fmt.Scanln(&input)
	if len(input) == 0 && def != "" {
		input = def
	}
	return input
}
