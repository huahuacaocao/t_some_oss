/*
@Time : 01/06/2020 4:58 PM 
@Author : GC
*/
package download

import (
	"fmt"
	"github.com/huahuacaocao/t_some_oss/config"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
	"time"
)

// 下载时资源标识使用 domain + url

// 公开空间访问 token
func PubToken() string {
	domain := "https://image.example.com"
	key := "测试.jpg"
	pubAccessURL := storage.MakePublicURL(domain, key)
	fmt.Println(pubAccessURL, len(pubAccessURL))
	return pubAccessURL
}

// 私有空间访问 token
// 地址格式 ?e=1451491200&token=MY_ACCESS_KEY:yN9WtB0lQheegAwva64yBuH3ZgU='
// token 中 MY_ACCESS_KEY 明文传输
func PriToken() string {
	credential := auth.New(config.ACCESS_KEY, config.SECRET_KEY)
	domain := "https://image.example.com"
	key := "测试.jpg"
	deadline := time.Now().Add(time.Second * 3600).Unix()
	privateAccessURL := storage.MakePrivateURL(credential, domain, key, deadline)
	fmt.Println(privateAccessURL, len(privateAccessURL))
	return privateAccessURL
	//privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	//fmt.Println(privateAccessURL)
}
