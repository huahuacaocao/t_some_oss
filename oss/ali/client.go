/*
@Time : 04/06/2020 2:53 PM 
@Author : GC
*/
package ali

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/huahuacaocao/t_some_oss/config"
)

//  sts 临时用户客户端
func stsClient() (*oss.Client, error) {
	credential := CreateStsToken3()
	return oss.New(
		config.ALIYUN_OSS_ENDPOINT,
		credential.AccessKeyId,
		credential.AccessKeySecret,
		oss.SecurityToken(credential.SecurityToken),
		oss.Timeout(60, 120),
	)
}
