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

// 判断 bucket 是否存在
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

// 获取 bucket 地区
func BucketRegion() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}

	if region, err := client.GetBucketLocation("bkt-t-001"); err != nil {
		panic(err)
	} else {
		fmt.Println(region)
	}
}

// bucket 信息
func BucketInfo() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}

	if info, err := client.GetBucketInfo("bkt-t-004"); err != nil {
		panic(err)
	} else {
		fmt.Println(info.BucketInfo.CreationDate)
		fmt.Println(info.BucketInfo.Name)
		fmt.Println(info.BucketInfo.RedundancyType)
		fmt.Println(info.BucketInfo.ACL)
		fmt.Println(info.BucketInfo.ExtranetEndpoint)
		fmt.Println(info.BucketInfo.IntranetEndpoint)
		fmt.Println("Owner:", info.BucketInfo.Owner.XMLName, info.BucketInfo.Owner.DisplayName, info.BucketInfo.Owner.ID) // 主账号(不是 RAM 账号)
		//{ Owner} 1288789039330709 1288789039330709
		fmt.Println(info.BucketInfo.SseRule)
		fmt.Println(info.BucketInfo.Versioning)
		fmt.Println(info.BucketInfo.XMLName)
	}
}

// 修改 bucket 权限
func SetBucketACL() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}
	if err = client.SetBucketACL("bkt-t-002", oss.ACLPublicRead); err != nil {
		panic(err)
	}
}

// 设置 bucket
func SetBucketTag() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}
	tags := []oss.Tag{
		oss.Tag{Key: "role", Value: "student"},
		oss.Tag{Key: "viplevel", Value: "1"},
	}

	if err = client.SetBucketTagging("bkt-t-002", oss.Tagging{Tags: tags}); err != nil {
		panic(err)
	}
}

// 获取 tag
func GetBucketTag() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}

	if tagging, err := client.GetBucketTagging("bkt-t-002"); err != nil {
		panic(err)
	} else {
		fmt.Println(tagging.Tags)
	}
}

// 根据 tag 获取 bucket
func BucketListWithTag() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}

	if list, err := client.ListBuckets(oss.TagKey("viplevel"), oss.TagValue("1")); err != nil {
		panic(err)
	} else {
		for _, v := range list.Buckets {
			fmt.Println(v)
		}
	}
}

// 设置生命周期规则
func BuildLiftcycle() {
	client, err := stsClient()
	if err != nil {
		panic(err)
	}
	//rule1 := oss.BuildLifecycleRuleByDays("rule1", "del/", true, 1)
	rule1 := oss.BuildLifecycleRuleByDays("rule1", "foo/", true, 300)

	deleteMarker := true
	//
	expiration := oss.LifecycleExpiration{ExpiredObjectDeleteMarker: &deleteMarker}
	// 版本控制 =>转换
	versionExpiration := oss.LifecycleVersionTransition{NoncurrentDays: 1}

	rule2 := oss.LifecycleRule{
		ID:         "rule2",
		Status:     "Enabled",
		Expiration: &expiration,
		//NonVersionTransition: &versionExpiration,
	}
	fmt.Println(rule1, rule2, versionExpiration)
	rules := []oss.LifecycleRule{rule2,rule1}
	if err := client.SetBucketLifecycle("bkt-t-002", rules); err != nil {
		fmt.Println(err)
	}

}
