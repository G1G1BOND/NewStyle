package utils

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func DelImg(url string) error {
	//accessKey := conf.C.QiNiu.AccessKey
	//secretKey := conf.C.QiNiu.SecretKey
	//bucket := conf.C.QiNiu.Bucket
	//qiNiuServer := conf.C.QiNiu.QiNiuServer
	//鉴权
	mac := qbox.NewMac(AccessKey, SecretKey)
	//配置属性
	cfg := storage.Config{
		UseHTTPS: false,
		Zone:     &storage.ZoneHuadongZheJiang2,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	//从url中解析出key并删除
	return bucketManager.Delete(Bucket, url[len(QiniuServer):])
}
