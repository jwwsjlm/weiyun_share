package weiyun

import (
	"github.com/microcosm-cc/bluemonday"
	"net/http"
)

func NewWeiyun(k string) *Weiyun {
	p := bluemonday.StrictPolicy()
	//var header http.Header
	header := http.Header{}
	header.Set("cookie", "wyctoken=0")
	header.Set("content-type", "application/json;charset=UTF-8")

	return &Weiyun{
		Key:        k,
		json:       `{"req_header":"{\"seq\":#seq,\"type\":1,\"cmd\":12002,\"appid\":30113,\"version\":3,\"major_version\":3,\"minor_version\":3,\"fix_version\":3,\"wx_openid\":\"\",\"user_flag\":0}","req_body":"{\"ReqMsg_body\":{\"ext_req_head\":{\"token_info\":{\"token_type\":0,\"login_key_type\":27,\"login_key_value\":\"\"},\"language_info\":{\"language_type\":2052}},\".weiyun.WeiyunShareViewMsgReq_body\":{\"share_pwd\":\"\",\"share_key\":\"#SHARE_KEY\"}}}"}`,
		client:     &http.Client{},
		bluemonday: p,
		header:     header,
	}
}
