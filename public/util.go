package public

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
)

// GenSaltPassword 生成通用加盐密码
func GenSaltPassword(salt, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	// %x 转换成16进制
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	return fmt.Sprintf("%x", s2.Sum(nil))
}

//MD5 md5加密
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Obj2Json 转换成JSON
func Obj2Json(s interface{}) string {
	bts, _ := json.Marshal(s)
	return string(bts)
}

// InStringSlice 查找slice中是否有str
func InStringSlice(slice []string,str string) bool{
	for  _,item:=range slice{
		if str==item{
			return true
		}
	}
	return false
}