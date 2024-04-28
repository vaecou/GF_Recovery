package validate

import (
	"net"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
)

// 判断元素是否存在于切片中
func InSlice[K comparable](slice []K, key K) bool {
	for _, v := range slice {
		if v == key {
			return true
		}
	}
	return false
}

// 判断IP是否为ipv4
func IsIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

// 检查密码
func CheckPassword(password, salt, hash string) (err error) {
	if hash != gmd5.MustEncryptString(password+salt) {
		err = gerror.New("密码不正确")
		return
	}
	return
}
