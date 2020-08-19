/*
@Time : 19/08/2020
@Author : GC
@Desc : 
*/

package token

import (
	"github.com/huahuacaocao/t_some_oss/config"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

type Qiniu struct {
	accessKey string
	secretKey string
}

var qiniu Qiniu

func getQiniu() *Qiniu {
	qiniu = Qiniu{
		accessKey: config.Conf.Qiniu.AccessKey,
		secretKey: config.Conf.Qiniu.SecretKey,
	}
	return &qiniu
}

//fmt.Printf("%+v", config.Conf.Qiniu.AccessKey)
// 简单 token
func (q *Qiniu) simpleToken() string {
	mac := qbox.NewMac(q.accessKey, q.secretKey)
	bkt := "bkt-t-1"
	putPolicy := storage.PutPolicy{
		Scope: bkt,
	}
	return putPolicy.UploadToken(mac)
}

// 带有返回信息的 token
// https://developer.qiniu.com/kodo/manual/1235/vars#magicvar
// 如果变量取值失败（例如在上传策略 (PutPolicy)中指定了一个并不存在的魔法变量），响应内容中对应的变量将被赋予空值
func (q *Qiniu) tokenWithReturnBody() string {
	mac := qbox.NewMac(q.accessKey, q.secretKey)
	bkt := "bkt-t-1"
	putPolicy := storage.PutPolicy{
		Scope: bkt,
		ReturnBody: `{"key":"${key}","hash":"${etag}","bucket":"${bucket}","etag":"${etag}","fname":"${fname}","fsize":"${fsize}",
"mimeType":"${mimeType}","endUser":"${endUser}","persistentId":"${persistentId}","ext":"${ext}","fprefix":"${fprefix}","uuid":"${uuid}",
"bodySha1":"${bodySha1}"
}`,
		//{$(imageInfo.width),
		//"imageInfo":"$(imageInfo)"
		//"exif":"${exif}"
		//","imageInfo":"${imageInfo}"
		//,,"avinfo":"${avinfo}"
		//,,"imageAve":"${imageAve}"

		//"year":"${year}",
		//"mon":"${mon}",
		//"day":"${day}",
		//"hour":"${hour}",
		//"min":"${min}",
		//"sec":"${sec}",
	}
	// TODO 指定 "imageInfo":"$(imageInfo)" 后会报错, 需要将内容逐个指出,  因 response 中 会双引号相互嵌套了 BUG
	return putPolicy.UploadToken(mac)
}
