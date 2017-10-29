package main

import (
	"os"
	"testing"
)

// test helper

func TestParseCsv(t *testing.T) {
	tmpFileName := "parse_test.csv"

	// setup
	file, err := os.Create(tmpFileName)
	if err != nil {
		t.Fatalf("テスト用ファイルの作成に失敗. err: %s", err)
	}
	input := "menu,price\nmelon,1000"
	file.Write(([]byte)(input))
	file.Close()

	fp, err := os.Open(tmpFileName)
	if err != nil {
		t.Fatalf("テスト用ファイルを開くのに失敗")
	}
	// test

	var actual [][]string = parseCsv(fp)
	if actual[0][0] != "menu" ||
		actual[0][1] != "price" ||
		actual[1][0] != "melon" ||
		actual[1][1] != "1000" {
		t.Errorf("Failed to parse. input=%s, actual=%s", input, actual)
	}

	// clean
	os.Remove(tmpFileName)
}
