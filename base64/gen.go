package main

import (
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("the numbers of args must be 1")
	}
	b := []byte(flag.Arg(0))

	sEnc := base64.StdEncoding.EncodeToString(b)
	fmt.Printf("enc=[%s]\n", sEnc)

	sDec, err := base64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		fmt.Printf("base64 decode failure, error=[%v]\n", err)
	} else {
		fmt.Printf("dec=[%s]\n", sDec)
	}
	fmt.Println("res:", cut(string(sEnc), 5))
}

func cut(src string, skip int) string {
	var ru = []rune(src)
	var res string
	for index, item := range ru {
		if index%skip == 0 {
			res += string(item)
		}
	}
	return res
}
