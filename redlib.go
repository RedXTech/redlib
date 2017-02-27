package redlib

import (
	//"fmt"
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"errors"
)

func Readln() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	input = strings.Replace(input, "\n", "", -1)
	return input
}

func Get(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New(err.Error())
	}
	defer resp.Body.Close()
	bs, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return "", errors.New(err2.Error())
	}
	return string(bs), nil
}
