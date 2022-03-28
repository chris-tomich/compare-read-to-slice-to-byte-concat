package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	Md5FileUsingReadToSlice()
	Md5FileUsingByteConcat()
}

func Md5FileUsingReadToSlice() {
	absImagePath, err := filepath.Abs("battle-of-fort-henry-1.jpg")

	if err != nil {
		panic(err)
	}

	image, err := os.Open(absImagePath)

	if err != nil {
		panic(err)
	}

	defer image.Close()
	imageFileInfo, err := image.Stat()

	if err != nil {
		panic(err)
	}

	imageNameBytes := []byte("battle-of-fort-henry-1.jpg")
	imageFrame := make([]byte, int64(len(imageNameBytes))+imageFileInfo.Size())
	imageNameSlice := imageFrame[:len(imageNameBytes)]
	imageDataSlice := imageFrame[len(imageNameBytes):]
	copy(imageNameSlice, imageNameBytes)

	bytesRead, err := image.Read(imageDataSlice)

	if err != nil {
		panic(err)
	}

	if int64(bytesRead) != imageFileInfo.Size() {
		fmt.Printf("Expected to read %v bytes, but instead read %v bytes.\n", imageFileInfo.Size(), bytesRead)
	}

	_ = md5.Sum(imageFrame)
	//fmt.Printf("%x\n", checksum)

	// raw, err := os.Create("read_to_slice.bin")

	// if err != nil {
	// 	panic(err)
	// }

	// defer raw.Close()
	// raw.Write(imageFrame)
}

func Md5FileUsingByteConcat() {
	absImagePath, err := filepath.Abs("battle-of-fort-henry-1.jpg")

	if err != nil {
		panic(err)
	}

	image, err := os.Open(absImagePath)

	if err != nil {
		panic(err)
	}

	defer image.Close()
	imageFileBytes, err := io.ReadAll(image)

	if err != nil {
		panic(err)
	}

	imageNameBytes := []byte("battle-of-fort-henry-1.jpg")
	imageFrame := append(imageNameBytes, imageFileBytes...)

	_ = md5.Sum(imageFrame)
	//fmt.Printf("%x\n", checksum)

	// raw, err := os.Create("byte_concat.bin")

	// if err != nil {
	// 	panic(err)
	// }

	// defer raw.Close()
	// raw.Write(imageFrame)
}

func Md5FileUsingByteConcatReversed() {
	absImagePath, err := filepath.Abs("battle-of-fort-henry-1.jpg")

	if err != nil {
		panic(err)
	}

	image, err := os.Open(absImagePath)

	if err != nil {
		panic(err)
	}

	defer image.Close()
	imageFileBytes, err := io.ReadAll(image)

	if err != nil {
		panic(err)
	}

	imageFrame := append(imageFileBytes, []byte("battle-of-fort-henry-1.jpg")...)

	_ = md5.Sum(imageFrame)
	//fmt.Printf("%x\n", checksum)
}
