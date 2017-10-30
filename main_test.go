package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseCsv(t *testing.T) {
	input := strings.NewReader("menu,price\nmelon,1000")
	var actual [][]string
	actual = parseCsv(input)
	if actual[0][0] != "menu" ||
		actual[0][1] != "price" ||
		actual[1][0] != "melon" ||
		actual[1][1] != "1000" {
		t.Errorf("Failed to parse. input=%s, actual=%s", input, actual)
	}
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

func TestInsertBar(t *testing.T) {
	input := []string{"menu", "price"}
	var actual string
	actual = insertBar(input)

	if actual != "|menu|price|" {
		t.Errorf("縦棒の挿入が間違っている. expected=\"|menu|price|\", actual=\"%s\"", actual)
	}
}

func TestPadWord(t *testing.T) {
	var actual string
	actual = padWord("word", 9)

	// 1. wordの後ろに、最大文字列のサイズになるまで空白でパディング
	// 2. 上の文字列の前後に空白をひとつずついれる
	expected := " word      "

	if actual != expected {
		t.Errorf("パディングが間違っている. expected=\"%s\", actual=\"%s\"", expected, actual)
	}
}

func TestGetMDMatrixLine(t *testing.T) {
	inputWords := []string{"menu", "price"}
	maxWordSizes := []int{5, 9}

	var actual string
	actual = getMDMatrixLine(inputWords, maxWordSizes)
	expected := "| menu  | price     |"

	if actual != expected {
		t.Errorf("マークダウンの行生成が間違っている. expected=\"%s\", actual=\"%s\"", expected, actual)
	}
}

func TestGetMDM2ndLine_Left(t *testing.T) {

	matrixType := "left"
	maxWordSizes := []int{3, 9}

	var actual string
	actual = getMDM2ndLine(matrixType, maxWordSizes)
	expected := "| :-- | :--       |"

	if actual != expected {
		t.Errorf("マークダウンの表の二行目の出力が間違っている. matrixType=\"%s\", expected=\"%s\", actual=\"%s\"", matrixType, expected, actual)
	}
}

func TestGetMDM2ndLine_Right(t *testing.T) {

	matrixType := "right"
	maxWordSizes := []int{3, 9}

	var actual string
	actual = getMDM2ndLine(matrixType, maxWordSizes)
	expected := "| --: | --:       |"

	if actual != expected {
		t.Errorf("マークダウンの表の二行目の出力が間違っている. matrixType=\"%s\", expected=\"%s\", actual=\"%s\"", matrixType, expected, actual)
	}
}

func TestGetMDM2ndLine_Center(t *testing.T) {

	matrixType := "center"
	maxWordSizes := []int{3, 9}

	var actual string
	actual = getMDM2ndLine(matrixType, maxWordSizes)
	expected := "| :-: | :-:       |"

	if actual != expected {
		t.Errorf("マークダウンの表の二行目の出力が間違っている. matrixType=\"%s\", expected=\"%s\", actual=\"%s\"", matrixType, expected, actual)
	}
}

func TestGetMDMDataLines(t *testing.T) {
	input := [][]string{{"menu", "price"}, {"melon", "50000"}, {"apple", "100"}}
	maxWordSizes := []int{5, 5}
	var actual []string
	actual = getMDMDataLines(input, maxWordSizes)
	expected := []string{"| melon | 50000 |", "| apple | 100   |"}

	if actual[0] != expected[0] || actual[1] != expected[1] {
		t.Errorf("マークダウンのデータ行の出力が間違っている. expected=\"%s\", actual=\"%s\"", expected, actual)
	}
}

func Test_getMDMDataLines(t *testing.T) {
	type args struct {
		input        [][]string
		maxWordSizes []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "マークダウンのデータ行の出力が間違っている",
			args: args{
				input:        [][]string{{"menu", "price"}, {"melon", "50000"}, {"apple", "100"}},
				maxWordSizes: []int{5, 5},
			},
			want: []string{"| melon | 50000 |", "| apple | 100   |"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMDMDataLines(tt.args.input, tt.args.maxWordSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMDMDataLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
