/*
@Time : 18/08/2020
@Author : GC
@Desc : 
*/

package main

import (
	"fmt"
	_ "github.com/huahuacaocao/t_some_oss"
	"github.com/huahuacaocao/t_some_oss/config"
	"github.com/huahuacaocao/t_some_oss/services"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	// 加载配置文件
	loadConf()
	//res, err := services.ImTokenService.CreateSimpleUploadToken()
	//res, err := services.ImTokenService.CreateTokenWithReturnBody()
	//res, err := services.ImUploadService.FormUpload()
	//res, err := services.ImUploadService.FragmentUpload()
	res, err := services.ImUploadService.ResumeUpload()
	fmt.Printf("%+v\n %+v\n", res, err)
}

func loadConf() {
	bytes, err := ioutil.ReadFile("/Users/mafengwo/Documents/okay/github/t_some_oss/brun/config.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(bytes, &config.Conf); err != nil {
		panic(err)
	}
}
