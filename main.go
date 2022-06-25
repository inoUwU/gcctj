package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f := flag.CommandLine.String("e", "", "e")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(*f)

	file, err := os.Open("./test/15NIIGAT.CSV")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	// []stringなのでループする

	var isSkip bool = true

	for {
		row, err := r.Read() // csvを1行ずつ読み込む
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
