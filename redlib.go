package redlib

import (
	"bufio"
	//"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"errors"
)

func Readln() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New(err.Error())
	}
	input = strings.Replace(input, "\n", "", -1)
	return input, nil
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
