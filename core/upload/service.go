/*
@Time : 19/08/2020
@Author : GC
@Desc : 
*/

package upload

import (
	"context"
	"fmt"
	"github.com/huahuacaocao/t_some_oss/services"
	"github.com/qiniu/api.v7/v7/storage"
	"sync"
)

type uploadService struct {
}

var once sync.Once

func init() {
	once.Do(func() {
		services.ImUploadService = new(uploadService)
	})
}

func (uploadService) FormUpload() (dto services.UploadReturnDTO, err error) {
	localFile := "/Users/mafengwo/Documents/okay/github/t_some_oss/brun/1595472042681.jpg"
	token, _ := services.ImTokenService.CreateTokenWithReturnBody()
	//token, _ := services.ImTokenService.CreateSimpleUploadToken()
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	//ret := storage.PutRet{}
	type IInfo struct {
		ColorModel string `json:"colorModel"`
		Width      string `json:"width"`
		Format     string `json:"format"`
		Height     string `json:"height"`
		Size       string `json:"size"`
	}

	//"imageInfo":"{"colorModel":"ycbcr","format":"jpeg","height":28,"size":6045,"width":66}
	//DebugMode
	//client.DebugMode = true
	ret := struct {
		Key          string `json:"key"`
		Hash         string `json:"hash"`
		Bucket       string `json:"bucket"`
		Etag         string `json:"etag"`
		Fname        string `json:"fname"`
		Fsize        string `json:"fsize"`
		MimeType     string `json:"mimeType"`
		EndUser      string `json:"endUser"`
		PersistentId string `json:"persistentId"`
		Ext          string `json:"ext"`
		Fprefix      string `json:"fprefix"`
		Uuid         string `json:"uuid"`
		BodySha1     string `json:"bodySha1"`
		ImageInfo    IInfo  `json:"imageInfo"`
		//ImageInfo interface{} `json:"imageInfo"`
	}{}

	//,
	//"persistentId":"${persistentId}","ext":"${ext}","fprefix":"${fprefix}","uuid":"${uuid}",
	//	"bodySha1":"${bodySha1}","imageInfo":"$(imageInfo.width)

	err = formUploader.PutFile(context.Background(), &ret, token.Token, "000001", localFile, nil)
	fmt.Println(context.Background())
	if err != nil {
		return nil, err
	}
	return ret, nil
}
