package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadJson(fileName string) []byte {
	//opening a json file
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// read our opened xmlFile as a byte array.
	jsonText, _ := ioutil.ReadAll(jsonFile)
	return jsonText
}

func OverwriteJson(filename string, data []byte) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err)
	}

	n, err := io.WriteString(f, string(data))
	if err != nil {
		fmt.Println(n, err)
	}
}