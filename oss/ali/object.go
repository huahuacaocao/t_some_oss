/*
@Time : 04/06/2020 4:09 PM 
@Author : GC
*/
package ali

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"strconv"
	"strings"
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

// 删除文件
func DeleteSingle() {
	bucket, err := bucket("bkt-t-002")
	if err != nil {
		panic(err)
	}
	if err = bucket.DeleteObject("t.png"); err != nil {
		panic(err)
	}
}

// 追加
func Append() {
	bucket, err := bucket("bkt-t-002")
	if err != nil {
		panic(err)
	}
	var nextPos int64 = 0
	str := "0123456789"
	nextPos, err = bucket.AppendObject("append", strings.NewReader(str), nextPos)
	if err != nil {
		panic(err)
	}
	fmt.Println(nextPos)
	nextPos, err = bucket.AppendObject("append", strings.NewReader(str), nextPos)
	if err != nil {
		panic(err)
	}
	fmt.Println(nextPos)
}

// 断点续传
func BreakpointResume() {
	bucket, err := bucket("bkt-t-002")
	if err != nil {
		panic(err)
	}

	err = bucket.UploadFile("breakpoint.CHM", "Go.CHM", 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	if err != nil {
		panic(err)
	}
}

// 分片上传
func SpliceUpload() {
	bucket, err := bucket("bkt-t-002")

	chunks, err := oss.SplitFileByPartNum("Go.CHM", 10)
	if err != nil {
		panic(err)
	}
	fd, err := os.Open("Go.CHM")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	imur, err := bucket.InitiateMultipartUpload("Go.CHM")
	bucket.InitiateMultipartUpload("splite.CHM")
	parts := []oss.UploadPart{}
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, io.SeekStart)
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			panic(err)
		}
		parts = append(parts, part)
	}

	cmur, err := bucket.CompleteMultipartUpload(imur, parts)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmur)
}

// 断点续传下载
func BreakpointResumeDownload() {
	bucket, err := bucket("bkt-t-002")
	err = bucket.DownloadFile("breakpoint.CHM", "bb.CHM", 1024*100, oss.Routines(10), oss.Checkpoint(true, "llllll"))
	if err != nil {
		panic(err)
	}
}
