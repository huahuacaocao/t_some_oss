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
	credential.SecurityToken = "CAIS8wF1q6Ft5B2yfSjIr5fzM9Pf2Ydg4KSPYBGGo3EmYr5rm5/+2zz2IHxJenFgA+EWtvs2m29V6v8blqJ8QNpESUHCccp+/8z8RqA8rMyT1fau5Jko1beHewHKeTOZsebWZ+LmNqC/Ht6md1HDkAJq3LL+bk/Mdle5MJqP+/UFB5ZtKWveVzddA8pMLQZPsdITMWCrVcygKRn3mGHdfiEK00he8TolufTgkpzEsUuF1wWlkr4vyt6vcsT+Xa5FJ4xiVtq55utye5fa3TRYgxowr/sr0vQUo2aX5o7EXAYMv0zfKYvIt9R1Kkp+fbMq5i2oY0kQWJcagAEs1hQfXzj8eqJtnOa4u2vG1ZMks+0p559CnMfiS54DQaG3Roci4OtRQGd9+myB1FI8HPLjtSozjQzMIoFRzvAYCczXzy7sDfxixuHokhb34lPENA4UbNKJkbauS8cI0KaAmxZVv6dNgJuEkQhSVwdQbl3rBOjhYxQksSi58HtW5A=="
	credential.AccessKeyId = "STS.NTFxik4XAWfdb77Gqsn1GtZU9"
	credential.AccessKeySecret = "FGK2do95c1k7N6CRqXR8P3R2uCSGJrZvD7S6Lhk4TMjW"
	credential.Expiration = "2020-06-04T09:39:15Z"
	return credential
}
