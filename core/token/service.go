/*
@Time : 18/08/2020
@Author : GC
@Desc : 
*/

package token

import (
	"errors"
	"github.com/huahuacaocao/t_some_oss/services"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		services.ImTokenService = new(tokenService)
	})
}

type tokenService struct {
}

// 创建简单token
func (tokenService) CreateSimpleUploadToken() (dto services.SimpleUploadTokenDTO, err error) {
	token := getQiniu().simpleToken()
	if token != "" {
		dto.Token = token
		return
	} else {
		err = errors.New("gen qiniu token fail")
	}
	return
}

// 创建带有回传信息(作为 Response)的 token
func (tokenService) CreateTokenWithReturnBody() (dto services.SimpleUploadTokenDTO, err error) {
	token := getQiniu().tokenWithReturnBody()
	if token != "" {
		dto.Token = token
		return
	} else {
		err = errors.New("gen qiniu token fail")
	}
	return
}
