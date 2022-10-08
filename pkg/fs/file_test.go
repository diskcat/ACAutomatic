package fs

import (
	"fmt"
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	data, err := LoadConfig("../../config/config.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func TestInsert(t *testing.T) {
	err := Insert("ddd", "../../config/config.txt")
	if err != nil {
		log.Println("添加敏感词失败！")
	}
}

func TestDelete(t *testing.T) {
	err := Delete("ddd", "../../config/config.txt")
	if err != nil {
		log.Println("删除失败！")
	}
}
