package fs

import (
	"io"
	"log"
	"os"
	"strings"
)

func ReadAll(fn string) ([]byte, error) {
	fd, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	if data, err := io.ReadAll(fd); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func LoadConfig(fn string) ([]string, error) {
	data, err := ReadAll(fn)
	if err != nil {
		return nil, err
	}
	str := string(data)
	words := strings.Split(str, "\r\n")
	return words, nil
}

func Insert(line string, fn string) error {
	fd, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	if _, err := fd.WriteString("\r\n" + line); err != nil {
		panic(err)
	}
	log.Println("敏感词添加成功")
	return nil
}

func Delete(line string, fn string) error {
	data, err := LoadConfig(fn)
	if err != nil {
		return err
	}
	tmp := []string{}
	for _, words := range data {
		if words == line {
			continue
		}
		tmp = append(tmp, words)
	}
	fd, err := os.OpenFile(fn, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		return err
	}
	for i, line := range tmp {
		if _, err := fd.WriteString(line); err != nil {
			panic(err)
		}
		if i != len(tmp)-1 {
			if _, err := fd.WriteString("\r\n"); err != nil {
				panic(err)
			}
		}
	}
	log.Println("删除成功！！")
	return nil
}
