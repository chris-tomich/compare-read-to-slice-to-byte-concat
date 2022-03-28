package main

import "testing"

func BenchmarkMd5FileUsingReadToSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Md5FileUsingReadToSlice()
	}
}

func BenchmarkMd5FileUsingByteConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Md5FileUsingByteConcat()
	}
}

func BenchmarkMd5FileUsingByteConcatReversed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Md5FileUsingByteConcatReversed()
	}
}
