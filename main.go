package main

import (
	"bufio"
	"container/list"
	"encoding/csv"
	"fmt"
	"log"
	"os"
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
func parseCsv(fp *os.File) [][]string {
	reader := csv.NewReader(fp)
	// reader.Comma = '\t'
	// reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない！
	var records [][]string
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed: Cannot parse input strings. err=\"%s\"", err)
	}
	fmt.Println(records)
	return records
}

func main() {

	in := list.New()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		in.PushBack(s.Text())
	}

	for e := in.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	if s.Err() != nil {
		log.Fatal(s.Err())
	}

}
