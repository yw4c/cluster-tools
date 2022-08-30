package sdk

import (
	"cluster-tools/config"
	"io/ioutil"
	"net/http"
)

type IEgressAddress interface {
	Get() string
}

type EgressAddress struct {
}

func (e EgressAddress) Get() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.GetConfigInstance().EgressIpURL, nil)
	if err != nil {
		return ""
	}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return ""
	} else {
		return string(body)
	}
}

func NewEgressAddress() IEgressAddress {
	return &EgressAddress{}
}
