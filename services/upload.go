/*
@Time : 19/08/2020
@Author : GC
@Desc : 
*/

package services

var ImUploadService UploadService

type UploadService interface {
	// 服务端表单上传
	FormUpload() (UploadReturnDTO, error)
}

type UploadReturnDTO interface {
}
