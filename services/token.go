/*
@Time : 18/08/2020
@Author : GC
@Desc : 
*/

package services

var ImTokenService TokenService

type TokenService interface {
	// 创建简单上传 token
	CreateSimpleUploadToken() (SimpleUploadTokenDTO, error)
	CreateTokenWithReturnBody() (SimpleUploadTokenDTO, error)
}

type SimpleUploadTokenDTO struct {
	Token string
}
