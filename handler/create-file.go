package handler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	fileName = "./coffee-product.ts"
)

func CreateFileBufio(data []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file failed, err:", err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return err
	}
	if n != len(data) {
		return errors.New("长度不够")
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("flush failed, err:", err)
		return err
	}
	return nil
}

func CreateFileIo(data []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("create file failed, err:", err)
		return err
	}
	defer file.Close()
	n, err := file.Write(data)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return err
	}
	if n != len(data) {
		fmt.Println("长度不够, n:", n)
		return errors.New("长度不够")
	}
	return nil
}
