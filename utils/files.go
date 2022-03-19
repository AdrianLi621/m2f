package files

import (
	"fmt"
	"os"
	"path/filepath"
)

const Tpl = "layout/layer.tpl"

/**
判断文件是否存在
*/
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

/**
递归创建文件
*/
func MkdirAll(filename string) bool {
	if !FileExist(filename) {
		d, f := filepath.Split(filename)
		if d != "" {
			err := os.MkdirAll(d, os.ModePerm)
			if err != nil {
				return false
			}
		}
		if f != "" {
			_, err := os.Create(filename)
			if err != nil {
				return false
			}
		}
	}
	return true
}

/**
写入文件
*/
func WriteToFile(filename string, bytes []byte) (bool, error) {
	if MkdirAll(filename) {
		err := os.WriteFile(filename, bytes, os.ModePerm)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
/**
生成model文件
 */
func GenerateStructs(filename, packages,params string) (bool, error) {
	bytes, err := os.ReadFile(Tpl)
	if err != nil {
		return false, err
	}
	b :=[]byte(fmt.Sprintf(string(bytes),StrFirstToUpper(packages),StrFirstToUpper(packages),params))
	if ok, err := WriteToFile(filename, b); !ok {
		return false, err
	}
	return true, nil
}


