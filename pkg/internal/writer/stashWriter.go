package writer

import (
	"bufio"
	"d2-item-sorter/pkg/domain"
	"fmt"
	"os"
	"time"
)

func WriteStashToFile(s domain.Stash, backupSuffix string) error {

	dataToWrite, itemCount := PrepareDataToWrite(s)
	err := backupAndWrite(s.Filepath, dataToWrite, backupSuffix)

	fmt.Printf("wrote %d items to: %s\n", itemCount, s.Filepath)
	return err
}

func backupAndWrite(path string, bytesToWrite []byte, backupSuffix string) error {
	if len(backupSuffix) == 0 {
		backupSuffix = ".backup_" + time.Now().Format("20060102_150405_000")
	}

	os.Rename(path, path+backupSuffix)

	file, err := os.Create(path)

	if err != nil {
		fmt.Printf("Error while opening file to write: %s", path)
		return err
	}

	defer file.Close()

	bfw := bufio.NewWriter(file)

	_, err = bfw.Write(bytesToWrite)

	bfw.Flush()

	return nil
}
