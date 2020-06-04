/*
@Time : 04/06/2020 2:47 PM 
@Author : GC
*/
package ali

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 创建 bucket(sts 用户)
func CreateBucket() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}
	err = client.CreateBucket("bkt-t-004", oss.ACL(oss.ACLPublicRead))
	// option
	// oss.StorageClass(oss.StorageArchive) 存储类型
	// oss.RedundancyType(oss.RedundancyZRS) 冗余方案

	if err != nil {
		panic(err)
	}
}

func bucket(bucketName string) (*oss.Bucket, error) {
	client, err := stsClient()
	if err != nil {
		return nil, err
	}
	return client.Bucket(bucketName)
}

func BucketList() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}
	marker := ""
	for {
		list, err := client.ListBuckets(oss.Marker(marker), oss.MaxKeys(2))
		// options
		// oss.Prefix()
		// oss.MaxKeys(2)

		if err != nil {
			panic(err)
		}
		for _, v := range list.Buckets {
			fmt.Println(v.Name, v.CreationDate)
		}
		if list.IsTruncated {
			marker = list.NextMarker
			// 这次循环最后一个 bucket 名
			fmt.Println(marker)
		} else {
			break
		}
	}
}

func ExistBucket() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}
	exist, err := client.IsBucketExist("bkt-t-001")
	if err != nil {
		panic(err)
	}
	fmt.Println(exist)
}
