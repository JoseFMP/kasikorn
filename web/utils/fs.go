package utils

import (
	"encoding/binary"
	"fmt"
	"os"
)

func DumpToFile(load []byte) error {

	logFile, errOpeningLogFile := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0755)
	if errOpeningLogFile != nil {
		return errOpeningLogFile
	}
	defer logFile.Close()

	err := binary.Write(logFile, binary.LittleEndian, load)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//	fmt.Printf("% x", buf.Bytes())
	return nil
}
