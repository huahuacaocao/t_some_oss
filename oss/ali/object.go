/*
@Time : 04/06/2020 4:09 PM 
@Author : GC
*/
package ali

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strconv"
	"time"
)

// 上传本地文件
func UploadSimple() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}
	//bkt-t-002
	oName := "t.png"
	localFile := "t.png"
	bucket, err := client.Bucket("bkt-t-002")
	if err != nil {
		panic(err)
	}
	err = bucket.PutObjectFromFile(oName, localFile)
	if err != nil {
		panic(err)
	}
}

// 流式上传
func UploadStream() {
	client, err := stsClient()
	bucket, err := client.Bucket("bkt-t-002")
	if err != nil {
		panic(err)
	}
	fd, err := os.Open("t.png")
	if err != nil {
		panic(err)
	}

	defer fd.Close()
	err = bucket.PutObject(strconv.FormatInt(time.Now().Unix(), 10)+"t.png", fd)
	if err != nil {
		panic(err)
	}
}

// 下载文件
func DownloadToFile() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}
	bucket, err := client.Bucket("bkt-t-002")
	if err != nil {
		panic(err)
	}
	oName := "1591257388t.png"
	saveFile := strconv.FormatInt(time.Now().Unix(), 10) + "tt.png"
	err = bucket.GetObjectToFile(oName, saveFile)
	if err != nil {
		panic(err)
	}
}

// 遍历文件列表
func ObjectList() {
	bucket, err := bucket("bkt-t-002")
	if err != nil {
		panic(err)
	}
	market := ""

	for {
		list, err := bucket.ListObjects(oss.Marker(market))
		if err != nil {
			panic(err)
		}
		for _, v := range list.Objects {
			fmt.Println(v.Key, v.ETag)
		}
		if list.IsTruncated {
			market = list.NextMarker
			fmt.Println(market)
		} else {
			// 全都返回了
			break
		}
	}
}

func DeleteSingle() {
	bucket, err := bucket("bkt-t-002")
	if err != nil {
		panic(err)
	}
	if err = bucket.DeleteObject("t.png"); err != nil {
		panic(err)
	}
}
