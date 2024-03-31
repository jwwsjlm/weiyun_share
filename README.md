# 腾讯微云笔记解析

```
https://share.weiyun.com/*****
```

传入链接末尾的字符串.代表k值也就是id

---



用法:

```
go get -u github.com/jwwsjlm/weiyun_share
```

```golang
	w := NewWeiyun("*******")
		json, err := w.WeiyunJson()//如果笔记内容为json内容的话.可以使用weiyunjson方法来解析.会过滤html标签内容.返回可被解析的json内容.自行解析
		if err != nil {
			panic(err)
		}
		fmt.Println(json)
```

具体可看weiyun_text.go文件具体例子

---



具体用到的代码:

[gjson ](https://github.com/tidwall/gjson "gjson")用于json解析

[bluemonday ](https://github.com/microcosm-cc/bluemonday "bluemonday ")用于html标签去除

---



下一版本增加密码.先写个例子自己玩玩
