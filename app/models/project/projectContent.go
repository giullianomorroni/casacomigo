package project

import (
	"fmt"
	"io/ioutil"
)

func ReadContent(file string) (content string, err error) {
    b, err := ioutil.ReadFile(file)
    if err != nil { fmt.Println(err) }
    return string(b), err
}