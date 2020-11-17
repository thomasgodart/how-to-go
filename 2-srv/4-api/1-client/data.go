package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// database structure

type DBDoc struct {
	ID        int
	CreatedAt *string
	UpdatedAt *string
	DeletedAt *string
	Name      string
	Content   string
}

type Json struct {
	All     []*DBDoc
	Get     DBDoc
	Set     DBDoc
	New     DBDoc
	Len     int
	Success bool
	Error   string
}

// api connection

var api Api

type Api struct {
}

func (api *Api) All() ([]*DBDoc, error) {
	var err error
	var dbDocs []*DBDoc

	res, err0 := http.PostForm("http://how-srv-api0:8080/all", url.Values{})
	if err0 != nil {
		return dbDocs, err0
	}

	bin, err1 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err1 != nil {
		return dbDocs, err1
	}

	var obj Json
	err2 := json.Unmarshal(bin, &obj)
	if err2 != nil {
		return dbDocs, err2
	}

	return obj.All, err
}

func (api *Api) Get(name string) (*DBDoc, error) {
	var err error
	var dbDoc DBDoc

	res, err0 := http.PostForm(fmt.Sprintf("http://how-srv-api0:8080/get/%s", name), url.Values{})
	if err0 != nil {
		return &dbDoc, err0
	}

	bin, err1 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err1 != nil {
		return &dbDoc, err1
	}

	var obj Json
	err2 := json.Unmarshal(bin, &obj)
	if err2 != nil {
		return &dbDoc, err2
	}

	return &obj.Get, err
}

func (api *Api) Set(name string, attr Obj) (*DBDoc, error) {
	var err error
	var dbDoc DBDoc

	res, err0 := http.PostForm(fmt.Sprintf("http://how-srv-api0:8080/set/%s", name), url.Values{"content": {attr["content"].(string)}})
	if err0 != nil {
		return &dbDoc, err0
	}

	bin, err1 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err1 != nil {
		return &dbDoc, err1
	}

	var obj Json
	err2 := json.Unmarshal(bin, &obj)
	if err2 != nil {
		return &dbDoc, err2
	}

	return &obj.Set, err
}

func (api *Api) Del(name string) (bool, error) {
	var err error

	res, err0 := http.PostForm(fmt.Sprintf("http://how-srv-api0:8080/del/%s", name), url.Values{})
	if err0 != nil {
		return false, err0
	}

	bin, err1 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err1 != nil {
		return false, err1
	}

	var obj Json
	err2 := json.Unmarshal(bin, &obj)
	if err2 != nil {
		return false, err2
	}

	if obj.Success != true {
		return false, err
	}

	return true, err
}

func (api *Api) New(name string, attr Obj) (*DBDoc, error) {
	var err error
	var dbDoc DBDoc

	res, err0 := http.PostForm(fmt.Sprintf("http://how-srv-api0:8080/new/%s", name), url.Values{"content": {attr["content"].(string)}})
	if err0 != nil {
		return &dbDoc, err0
	}

	bin, err1 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err1 != nil {
		return &dbDoc, err1
	}

	var obj Json
	err2 := json.Unmarshal(bin, &obj)
	if err2 != nil {
		return &dbDoc, err2
	}

	return &obj.New, err
}
