package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// やりかた
// 入力フェイズ
// 1. 入力を2次元配列に入れる。(csvのパース)
// 2. 各列の最大文字数を計算し、計算結果をいれた配列をつくる。
//
// 出力フェイズ
// 1. 1行目を出力する。
//    頭と末尾に|を入れる。配列[0][*]の各ワード間に|を入れる。ワードは、最大文字列の長さになるように前後に空白を入れる
// 2. 2行目を出力する。
//    書くワードを:始まりで、最大文字列になるように-を入れる
// 3. 3行目以降を出力する。
//    配列[1][*]以降を1行目と同じフォーマットで出力する

//
// 入力フェイズの関数
//

// csvを2次元配列にして返す
func parseCsv(r io.Reader) [][]string {
	reader := csv.NewReader(r)
	// reader.Comma = '\t'
	// reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない！
	var records [][]string
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed: Cannot parse input strings. err=\"%s\"", err)
	}
	return records
}

func getMaxWordSizes(matrix [][]string) []int {
	numberOfCol := len(matrix[0])
	maxWordSizes := make([]int, numberOfCol)
	for i := 0; i < numberOfCol; i++ {
		col := getCol(matrix, i)
		maxWordSizes[i] = getMaxWordSize(col)
	}
	return maxWordSizes
}

func getCol(matrix [][]string, index int) []string {
	numberOfRow := len(matrix)
	column := make([]string, numberOfRow)
	for i := 0; i < numberOfRow; i++ {
		column[i] = matrix[i][index]
	}
	return column
}

func getMaxWordSize(words []string) int {
	max := 0
	for i := 0; i < len(words); i++ {
		if len(words[i]) > max {
			max = len(words[i])
		}
	}
	return max
}

//
// 出力フェイズの関数
//
func insertBar(words []string) string {
	inserted := strings.Join(words, "|")
	return "|" + inserted + "|"
}

func padWord(word string, size int) string {
	format := " %-" + strconv.Itoa(size) + "s "
	return fmt.Sprintf(format, word)
}

func getMDMatrixLine(inputWords []string, maxWordSizes []int) string {
	var padedWords []string
	for i, word := range inputWords {
		padedWords = append(padedWords, padWord(word, maxWordSizes[i]))
	}
	return insertBar(padedWords)
}

func main() {

	//標準入力の読み込み
	var inputStrings []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		inputStrings = append(inputStrings, s.Text())
	}

	if s.Err() != nil {
		log.Fatal(s.Err())
	}

	inputString := strings.Join(inputStrings, "\n")

	//csvパース
	matrix := parseCsv(strings.NewReader(inputString))

	//各列の最大文字数
	maxWordSizes := getMaxWordSizes(matrix)

	fmt.Printf("maxWordSizes: %s", maxWordSizes)
}