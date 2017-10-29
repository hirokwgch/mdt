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
	var actual [][]string
	actual = parseCsv(fp)
	if actual[0][0] != "menu" ||
		actual[0][1] != "price" ||
		actual[1][0] != "melon" ||
		actual[1][1] != "1000" {
		t.Errorf("Failed to parse. input=%s, actual=%s", input, actual)
	}

	// clean
	os.Remove(tmpFileName)
}

func TestGetMaxWordSizes(t *testing.T) {
	input := [][]string{{"menu", "price"}, {"melon", "123456789"}}
	var actual []int
	actual = getMaxWordSizes(input)

	if actual[0] != 5 || actual[1] != 9 {
		t.Errorf("各列の最大サイズが間違っている. expected=[5 9], actual=[%d %d]", actual[0], actual[1])
	}
}

func TestGetCol(t *testing.T) {
	input := [][]string{{"menu", "price"}, {"melon", "123456789"}}
	var col0, col1 []string
	col0 = getCol(input, 0)

	if col0[0] != "menu" || col0[1] != "melon" {
		t.Errorf("間違った列データを返している. expected=[menu melon], actual=[%s %s]", col0[0], col0[1])
	}
	col1 = getCol(input, 1)
	if col1[0] != "price" || col1[1] != "123456789" {
		t.Errorf("間違った列データを返している. expected=[price 123456789], actual=[%s %s]", col1[0], col1[1])
	}
}

func TestGetMaxWordSize(t *testing.T) {
	input := []string{"price", "1234567890"}
	var actual int
	actual = getMaxWordSize(input)

	if actual != 10 {
		t.Errorf("最大サイズが間違っている. expected=10, actual=%d", actual)
	}
}
