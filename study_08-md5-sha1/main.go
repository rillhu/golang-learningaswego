package main

import (
	"crypto/md5"
	"fmt"
	"crypto/sha1"
	"os"
	"io"
)

func main() {

	testString := "Hi,pandaman"

	md5Inst := md5.New()
	md5Inst.Write([]byte(testString)) //Input need slice type
	res := md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", res)

	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(testString))
	res = sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", res)

	testFile := "123.txt"
	inFile, inerr := os.Open(testFile)

	if inerr == nil {
		md5f := md5.New()
		io.Copy(md5f,inFile)
		fmt.Printf("%x %s\n\n", md5.Sum([]byte("")), testFile)

		sha1f := sha1.New()
		io.Copy(sha1f, inFile)
		fmt.Printf("%x %s\n", sha1.Sum([]byte("")), testFile)
	} else {
		fmt.Println(inerr)
		os.Exit(1)
	}

}
