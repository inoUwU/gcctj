package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// TODO
// 動的な構造体の作成
// Jsonファイル出力
// コマンドライン引数の読み込み解析
// 実行時のパス取得

func main() {
	// 引数は後で実装
	// f := flag.CommandLine.String("e", "", "e")
	// flag.Parse()
	// for i, v := range os.Args {
	// 	fmt.Printf("args[%d] -> %s\n", i, v)
	// }
	// fmt.Println(*f)

	csvFile, err := os.Open("./test/15NIIGAT.CSV")

	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	firstRow, errFirstRow := r.Read()
	if errFirstRow != nil {
		log.Fatal(errFirstRow)
	}
	println(firstRow)

	// []stringなのでループする
	var isSkip bool = true
	for {
		// csvを1行ずつ読み込む
		row, err := r.Read()
		if isSkip {
			// 1行目はスキップする
			isSkip = false
			continue
		}

		if err == io.EOF {
			break // 読み込むべきデータが残っていない場合，Readはnil, io.EOFを返すため、ここでループを抜ける
		}

		// 配列から取り出す
		var joined string
		for i := 0; i < len(row); i++ {
			joined = joined + row[i]
		}

		//fmt.Println(joined)
	}
}
