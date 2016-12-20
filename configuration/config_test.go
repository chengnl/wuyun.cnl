package configuration

import (
	"fmt"
	"testing"
)

func TestInitFile(t *testing.T) {
	InitConfigFile("./config_test_file.conf")
	val, err := GetBool("bool")
	if err != nil {
		fmt.Println("GetBool err=%v\n", err)
	}
	fmt.Printf("GetBool=%v\n", val)

	val, err = GetBoolDefaultVal("bool", true)
	if err != nil {
		fmt.Println("GetBool err=%v\n", err)
	}
	fmt.Printf("GetBool=%v\n", val)

	intVal, intErr := GetInt("int8", 10, 8)
	if intErr != nil {
		fmt.Println("GetInt err=%v\n", intErr)
	}
	fmt.Printf("GetInt=%v\n", intVal)
}
