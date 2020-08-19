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
	// 分片上传
	FragmentUpload() (UploadReturnDTO, error)
	// 分片上传
	ResumeUpload() (UploadReturnDTO, error)
}

type UploadReturnDTO interface {
}
