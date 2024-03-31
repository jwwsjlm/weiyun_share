package weiyun

import (
	"errors"
	"github.com/tidwall/gjson"
	"html"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"weiyun/utils"
)

func (w *Weiyun) WeiyunJson() (string, error) {

	u := "https://share.weiyun.com/webapp/json/weiyunShareNoLogin/WeiyunShareView?g_tk=0&_" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	j := strings.Replace(w.json, "#seq", strconv.FormatInt(time.Now().UnixMilli(), 10)+strconv.Itoa(utils.GetARandom(4)), -1)
	j = strings.Replace(j, "#SHARE_KEY", w.Key, -1)
	req, err := http.NewRequest("POST", u, strings.NewReader(j))
	if err != nil {
		return "", err
	}
	req.Header = w.header

	resp, err := w.client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if !gjson.Valid(string(bodyText)) {
		return string(bodyText), errors.New("invalid json")
	}

	value := gjson.Get(string(bodyText), "data.rsp_body.RspMsg_body.note_list.0.html_content")
	if !value.Exists() {
		return "", errors.New("not found")

	}

	escapeString := w.bluemonday.Sanitize(value.String())
	escapeString = html.UnescapeString(escapeString)

	return escapeString, nil
}
