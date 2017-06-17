package redlib

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"unsafe"
)

func Readln() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
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

func Post(postUrl string, data map[string]string) string {
	hc := http.Client{}
	routerURL := postUrl
	form := url.Values{}

	for k, v := range data {
		form.Add(k, v)
	}

	req, err := http.NewRequest("POST", routerURL, strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.PostForm = form
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64; rv:43.0) Gecko/20100101 Firefox/43.0")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	b := buf.Bytes()
	s := *(*string)(unsafe.Pointer(&b))
	return s
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	} else {
		HandleErr(err)
	}
	return true
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func HandleErr(err error) bool {
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		return true
	}
}
