package configuration

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var commentChars = []byte("#!")
var separators = []byte("=:")
var configs map[string]string

func InitConfigFile(filePath string) error {
	cfile, err := os.Open(filePath)
	if err != nil {
		panic(err)
		return err
	}
	configs = make(map[string]string)
	info := bufio.NewReader(cfile)
	for line, _, err := info.ReadLine(); err == nil; line, _, err = info.ReadLine() {
		if len(line) > 0 {
			if isComment(line[0]) {
				continue
			} else {
				isHaveSep, index := separatorIndex(line)
				if isHaveSep {
					key := line[:index]
					value := line[index+1:]
					if len(key) > 0 && len(value) > 0 {
						configs[string(key)] = string(value)
					}
				}
			}
		}
	}
	return nil
}

//获取key的值字符串，没有key值返回""
func GetString(key string) string {
	return configs[key]
}

//获取key的值字符串，没有key值返回defaultVal
func GetStringDefaultVal(key string, defaultVal string) string {
	if config, ok := configs[key]; ok {
		return config
	}
	return defaultVal
}

// base：进位制（2 进制到 36 进制）
// bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
func GetInt(key string, base int, bitSize int) (int64, error) {
	if config, ok := configs[key]; ok {
		return strconv.ParseInt(config, base, bitSize)
	}
	return -1, fmt.Errorf("not found key:%s", key)
}

// base：进位制（2 进制到 36 进制）
// bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
//defaultVal 缺省值，没有key时返回defaultVal
func GetIntDefaultVal(key string, base int, bitSize int, defaultVal int64) (int64, error) {
	if config, ok := configs[key]; ok {
		return strconv.ParseInt(config, base, bitSize)
	}
	return defaultVal, nil
}

// base：进位制（2 进制到 36 进制）
// bitSize：指定整数类型（0:uint、8:uint8、16:uint16、32:uint32、64:uint64）
// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
//defaultVal 缺省值，没有key时返回defaultVal
func GetUint(key string, base int, bitSize int) (uint64, error) {
	if config, ok := configs[key]; ok {
		return strconv.ParseUint(config, base, bitSize)
	}
	return 0, fmt.Errorf("not found key:%s", key)
}

// base：进位制（2 进制到 36 进制）
// bitSize：指定整数类型（0:uint、8:uint8、16:uint16、32:uint32、64:uint64）
// 如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
//defaultVal 缺省值，没有key时返回defaultVal
func GetUintDefaultVal(key string, base int, bitSize int, defaultVal uint64) (uint64, error) {
	if config, ok := configs[key]; ok {
		return strconv.ParseUint(config, base, bitSize)
	}
	return defaultVal, nil
}
func GetBool(key string) (bool, error) {
	if config, ok := configs[key]; ok {
		return strconv.ParseBool(config)
	}
	return false, fmt.Errorf("not found key:%s", key)
}

//defaultVal 缺省值，没有key时返回defaultVal
func GetBoolDefaultVal(key string, defaultVal bool) (bool, error) {
	if config, ok := configs[key]; ok {
		return strconv.ParseBool(config)
	}
	return defaultVal, nil
}

// bitSize：指定整数类型（32:float32、64:float64）
func GetFloat(key string, bitSize int) (float64, error) {
	if config, ok := configs[key]; ok {
		return strconv.ParseFloat(config, bitSize)
	}
	return 0.0, fmt.Errorf("not found key:%s", key)
}

//bitSize：指定整数类型（32:float32、64:float64）
//defaultVal 缺省值，没有key时返回defaultVal
func GetFloatDefaultVal(key string, bitSize int, defaultVal float64) (float64, error) {
	if config, ok := configs[key]; ok {
		return strconv.ParseFloat(config, bitSize)
	}
	return defaultVal, nil
}
func isComment(preline byte) bool {
	for _, commetChar := range commentChars {
		if commetChar == preline {
			return true
		}
	}
	return false
}
func separatorIndex(line []byte) (isHaveSep bool, index int) {
	for index, char := range line {
		for _, sep := range separators {
			if char == sep {
				return true, index
			}
		}
	}
	return false, -1
}

func isFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
