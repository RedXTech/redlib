package redlib

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"errors"
	"net/url"
	"bytes"
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