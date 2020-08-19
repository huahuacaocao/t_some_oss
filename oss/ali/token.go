/*
@Time : 04/06/2020 2:46 PM 
@Author : GC
*/
package ali

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	aliyunsts "github.com/aliyun/aliyun-sts-go-sdk/sts"
	"github.com/huahuacaocao/t_some_oss/config"
	"strconv"
)

func CreateStsToken() sts.Credentials {
	request := sts.CreateAssumeRoleRequest()
	request.RoleSessionName = "gc-identify"
	request.Scheme = "https"
	//request.Policy = ""
	request.RoleArn = config.ALIYUN_ROLEARN
	request.DurationSeconds = requests.Integer(strconv.Itoa(config.STS_DURATIONSECONDS))

	client, err := sts.NewClientWithAccessKey(config.ALIYUN_OSS_REGIONID, config.ALIYUN_ACCESSKEY_ID, config.ALIYUN_ACCESSKEY_SECRET)
	if err != nil {
		panic(err)
	}
	resp, err := client.AssumeRole(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.GetHttpContentString())
	return resp.Credentials
}

func CreateStsToken2() aliyunsts.Credentials {
	client := aliyunsts.NewClient(config.ALIYUN_ACCESSKEY_ID, config.ALIYUN_ACCESSKEY_SECRET, config.ALIYUN_ROLEARN, "gc-identify")
	resp, err := client.AssumeRole(uint(config.STS_DURATIONSECONDS))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	return resp.Credentials
}

func CreateStsToken3() sts.Credentials {
	credential := sts.Credentials{}
	credential.SecurityToken = "CAIS8wF1q6Ft5B2yfSjIr5DfI/mN3+5jwbe4MHfr3FQEXcpVt6v/lzz2IHxJenFgA+EWtvs2m29V6v8blqJ8QNpESUHCccp+/8yiY/1Ko8yT1fau5Jko1beHewHKeTOZsebWZ+LmNqC/Ht6md1HDkAJq3LL+bk/Mdle5MJqP+/UFB5ZtKWveVzddA8pMLQZPsdITMWCrVcygKRn3mGHdfiEK00he8TolufTgkpzEsUuF1wWlkr4vyt6vcsT+Xa5FJ4xiVtq55utye5fa3TRYgxowr/sr0vQUo2aX5o7EXAYMv0zfKYvIt9R1Kkp+fbMq5i2oY0kQWJcagAEQlJB7fUUiGlb5Q5RG2K0EXxjNyDOBZ4lP4GAC9bNngm5DQP+kTShLo/uR1vj77M2J4Nmn2/CBJJKlRb1uDxJf5XPBHXBw4ZrQQJN6VmpuyljXAFbT2iRy0H/nt9G6o0FjX8f7+ky8fmo132443oE/bptci3G8ns6baeHGCzA7vg=="
	credential.AccessKeyId = "STS.NSjhC921BvuS2QZ8TQQEyXnTu"
	credential.AccessKeySecret = "HwRtXrgsY85hTMQ5qQcnK2ud34sN4XmN2QTCR5XFepZj"
	credential.Expiration = "2020-06-04T17:34:10Z"
	return credential
}
