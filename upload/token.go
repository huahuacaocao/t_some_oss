/*
@Time : 28/05/2020 1:29 PM 
@Author : GC
*/
package upload

import (
	"encoding/base64"
	"fmt"
	"github.com/huahuacaocao/t_some_oss/config"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
	"strings"
)

// 测试生成 uploadToken

var bucket string = ""

/**
storage.PutPolicy 策略结构体说明
https://developer.qiniu.com/kodo/manual/1206/put-policy

upToken 的值格式 AccessKey(明文) + ':' + encodedSign + ':' + encodedPutPolicy
即 AccessKey 在网络上可以任意传播
*/
func SimpleToken() string {
	putPolicy := storage.PutPolicy{}
	putPolicy.Scope = fmt.Sprintf("%s", "qiniu-ts-demo")
	credential := auth.New(config.ACCESS_KEY, config.SECRET_KEY)
	upToken := putPolicy.UploadToken(credential)
	fmt.Println(upToken)
	return upToken
}

// 带有有效期的Token
// 粗略测试, 500 毫秒之内生成 token 相同
func TimelineToken() string {
	putPolicy := storage.PutPolicy{}
	putPolicy.Scope = fmt.Sprintf("%s", "qiniu-ts-demo")
	putPolicy.Expires = 3600 // 1 个小时有效
	credential := auth.New(config.ACCESS_KEY, config.SECRET_KEY)
	upToken := putPolicy.UploadToken(credential)
	fmt.Println(upToken)
	return upToken
}

// 覆盖 token
// scope
// 		<bucket>: 只能新增
//		<bucket>:<key>: 默认允许修改(即覆盖)
//		<bucket>:<key> + insertOnly=1 : 只能新增
//		<bucket>:<keyPrefix> + isPrefixalScope=1 : 只能新增
func OverwriteToken() string {
	keyToOverwrite := "qiniu.mp4"
	pubPolicy := storage.PutPolicy{}
	pubPolicy.Scope = fmt.Sprintf("%s:%s", "qiniu-ts-demo", keyToOverwrite)
	credential := auth.New(config.ACCESS_KEY, config.SECRET_KEY)
	upToken := pubPolicy.UploadToken(credential)
	fmt.Println(upToken)
	return upToken
}

// 自定义上传回复凭证
// 这个 token 好长(306)
// returnBody 支持魔法变量(https://developer.qiniu.com/kodo/manual/1235/vars#magicvar)
// 和自定义变量(https://developer.qiniu.com/kodo/manual/1235/vars#xvar)
func CustomResponseToken() string {
	putPolicy := storage.PutPolicy{}
	putPolicy.Scope = fmt.Sprintf("%s", "qiniu-ts-demo")
	putPolicy.ReturnBody = `{"key":"$(key)","hash":"$(etag)","fsize":"$(fsize)","bucket":"$(bucket)","name":"$(x:name)"}`
	credential := auth.New(config.ACCESS_KEY, config.SECRET_KEY)
	upToken := putPolicy.UploadToken(credential)
	fmt.Println(upToken, len(upToken))
	return upToken
}

// 带有回调业务服务器的凭证(JSON)
func CallbackToken() string {
	putPolicy := storage.PutPolicy{
		Scope:            fmt.Sprintf("%s", "qiniu-ts-demo"),
		CallbackURL:      "http://api.example.com/qiniu/upload/callback",
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":"$(fsize)","bucket":"$(bucket)","name":"$(x:name)"}`,
		CallbackBodyType: "application/json",
	}
	credential := auth.New(config.ACCESS_KEY, config.SECRET_KEY)
	upToken := putPolicy.UploadToken(credential)
	fmt.Println(upToken, len(upToken))
	return upToken
}

// 带有回调业务服务器的凭证(URL)
func CallbackUrlToken() {
	putPolicy := storage.PutPolicy{
		Scope:        fmt.Sprintf("%s", "qiniu-ts-demo"),
		CallbackURL:  "http://api.example.com/qiniu/upload/callback",
		CallbackBody: "key=$(key)&hash=$(etag)&bucket=$(etag)&fsize=$(fsize)&name=$(x:name)",
	}
	credential := auth.New(config.ACCESS_KEY, config.SECRET_KEY)
	upToken := putPolicy.UploadToken(credential)
	fmt.Println(upToken, len(upToken))
}

// 带数据处理的凭证(即转码存储)
// 通过 persistentXXX 系列属性实现
// persistentOps: 资源上传成功后触发执行的预转持久化指令列表
// 		归档存储不支持
// persistentNotifyUrl: 接收持久化处理结果通知的 URL
//		json post
// persistentPipeline: 转码队列名
//		分公用队列和专用队列
func PersistentOpsToken() string {
	// persistentOps 预转数据处理命令 + 保存处理结果(指定空间和资源名)
	// persistentId 字段，唯一标识此任务
	// 结果存储: 默认使用上传时的 bucket, key由七牛生成
	// 指定存储结果 通过 | 拼接 saveas/EncodedEntryURI

	// 指定转码后的存储空间和名(存在覆盖风险)
	/**
	EncodedEntryURI格式
	entry = '<Bucket>:<Key>'
    encodedEntryURI = urlsafe_base64_encode(entry)
	*/
	saveMp4Entity := base64.URLEncoding.EncodeToString([]byte("qiniu-ts-demo" + ":avthumb_test_target.mp4"))
	saveJpgEntity := base64.URLEncoding.EncodeToString([]byte("qiniu-ts-demo" + ":vframe_test_targe.jpg"))

	// 指令
	// 转码格式说明 https://developer.qiniu.com/dora/api/1248/audio-and-video-transcoding-avthumb
	avthumbMp4Fop := "avthumb/mp4|saveas/" + saveMp4Entity
	vframeJpgFop := "vframe/jpg/offset/1|saveas/" + saveJpgEntity
	persistentOps := strings.Join([]string{avthumbMp4Fop, vframeJpgFop}, ";")
	// 执行队列名
	pipline := "test"
	putPolicy := storage.PutPolicy{
		Scope:               fmt.Sprintf("qiniu-ts-demo"),
		PersistentOps:       persistentOps,
		PersistentNotifyURL: "http://api.example.com/qiniu/pfop/notify",
		PersistentPipeline:  pipline,
	}
	credential := auth.New(config.ACCESS_KEY, config.SECRET_KEY)
	upToken := putPolicy.UploadToken(credential)
	fmt.Println(upToken, len(upToken))
	return upToken
}
