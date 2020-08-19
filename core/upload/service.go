/*
@Time : 19/08/2020
@Author : GC
@Desc : 
*/

package upload

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/huahuacaocao/t_some_oss/services"
	"github.com/qiniu/api.v7/v7/client"
	"github.com/qiniu/api.v7/v7/storage"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type uploadService struct {
}

var once sync.Once

func init() {
	once.Do(func() {
		services.ImUploadService = new(uploadService)
	})
}

func (uploadService) FormUpload() (dto services.UploadReturnDTO, err error) {
	localFile := "/Users/mafengwo/Documents/okay/github/t_some_oss/brun/1595472042681.jpg"
	token, _ := services.ImTokenService.CreateTokenWithReturnBody()
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	//ret := storage.PutRet{}
	type IInfo struct {
		ColorModel string `json:"colorModel"`
		Width      string `json:"width"`
		Format     string `json:"format"`
		Height     string `json:"height"`
		Size       string `json:"size"`
	}

	//"imageInfo":"{"colorModel":"ycbcr","format":"jpeg","height":28,"size":6045,"width":66}
	//DebugMode
	//client.DebugMode = true
	ret := struct {
		Key          string `json:"key"`
		Hash         string `json:"hash"`
		Bucket       string `json:"bucket"`
		Etag         string `json:"etag"`
		Fname        string `json:"fname"`
		Fsize        string `json:"fsize"`
		MimeType     string `json:"mimeType"`
		EndUser      string `json:"endUser"`
		PersistentId string `json:"persistentId"`
		Ext          string `json:"ext"`
		Fprefix      string `json:"fprefix"`
		Uuid         string `json:"uuid"`
		BodySha1     string `json:"bodySha1"`
		ImageInfo    IInfo  `json:"imageInfo"`
		//ImageInfo interface{} `json:"imageInfo"`
	}{}

	err = formUploader.PutFile(context.Background(), &ret, token.Token, "000001", localFile, nil)
	fmt.Println(context.Background())
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// 分片上传
func (uploadService) FragmentUpload() (dto services.UploadReturnDTO, err error) {
	localFile := "/Users/mafengwo/Documents/okay/github/t_some_oss/brun/树的基本概念】.mp4"
	client.DebugMode = true
	token, _ := services.ImTokenService.CreateTokenWithReturnBody()
	cfg := storage.Config{}
	resumeUploader := storage.NewResumeUploader(&cfg)
	var res interface{}
	err = resumeUploader.PutFile(context.Background(), &res, token.Token, "000002", localFile, nil)
	if err != nil {
		return
	}
	return res, nil
}

// 断点续传
func (uploadService) ResumeUpload() (dto services.UploadReturnDTO, err error) {
	localFile := "/Users/mafengwo/Documents/okay/github/t_some_oss/brun/树的基本概念】.mp4"
	token, _ := services.ImTokenService.CreateTokenWithReturnBody()
	cfg := storage.Config{}
	resumeUploader := storage.NewResumeUploader(&cfg)

	key := "000003"
	// 进度存储文件,唯一性 md5(bucket+key+local_path+local_file_last_modified)+".progress"
	fileInfo, err := os.Stat(localFile)
	if err != nil {
		return
	}
	fileSize := fileInfo.Size()
	recordPath, err := getProgressRecordFile(key, localFile, fileInfo)
	if err != nil {
		return
	}
	progressRecord := getprogressRecord(recordPath, fileSize)
	var res interface{}

	progressLock := sync.RWMutex{}
	putExtra := storage.RputExtra{
		Progresses: progressRecord.Progresses,
		Notify: func(blkIdx int, blkSize int, ret *storage.BlkputRet) {
			progressLock.Lock()
			progressLock.Unlock()
			//将进度序列化，然后写入文件
			progressRecord.Progresses[blkIdx] = *ret
			progressBytes, _ := json.Marshal(progressRecord)
			fmt.Println("write progress file", blkIdx, recordPath)
			wErr := ioutil.WriteFile(recordPath, progressBytes, 0644)
			if wErr != nil {
				fmt.Println("write progress file error,", wErr)
			}
		},
	}
	err = resumeUploader.PutFile(context.Background(), &res, token.Token, key, localFile, &putExtra)
	if err != nil {
		return
	}
	os.Remove(recordPath)
	return res, err
}

// 生成进度文件
func getProgressRecordFile(key string, localFile string, info os.FileInfo) (path string, err error) {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s:%s:%s:%s", "bkt-t-1", key, localFile, info.ModTime().UnixNano())))
	recordKey := hex.EncodeToString(h.Sum(nil)) + ".progress"
	recordDir := "/tmp/Temp/progress"
	if err = os.MkdirAll(recordDir, 0755); err != nil {
		return
	}
	path = filepath.Join(recordDir, recordKey)
	return
}

// 进度文件解析
func getprogressRecord(recordPath string, size int64) (record progressRecord) {
	tmpProcesses := make([]storage.BlkputRet, storage.BlockCount(size))
	// 尝试从旧的进度文件中读取进度
	file, err := os.Open(recordPath)
	if err != nil {
		record.Progresses = tmpProcesses
		return
	}

	progressBytes, err := ioutil.ReadAll(file)
	if err != nil {
		record.Progresses = tmpProcesses
		return
	}

	if err = json.Unmarshal(progressBytes, &record); err != nil {
		for _, item := range record.Progresses {
			if storage.IsContextExpired(item) {
				record.Progresses = tmpProcesses
				return
			}
		}
	}
	file.Close()
	if len(record.Progresses) == 0 {
		record.Progresses = tmpProcesses
	}
	return
}

type progressRecord struct {
	Progresses []storage.BlkputRet `json:"progresses"`
}
