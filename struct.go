package weiyun

import (
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

type Weiyun struct {
	Key        string
	json       string
	client     *http.Client
	bluemonday *bluemonday.Policy
	header     http.Header
}
