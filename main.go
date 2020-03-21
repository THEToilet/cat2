package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	filename string
)

func init() {
	flag.StringVar(&filename, "filename", "", "catしたいファイル名")
	flag.Parse()
}

func cat(file string) error {

	fi, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("ファイルをひらけませんでした。")
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	// sfから1行ずつ読み込み、"行数:"を前に付けて標準出力に書き出す。
	for i := 1; scanner.Scan(); i++ {
		fmt.Fprintf(os.Stdout, "%d:%s\n", i, scanner.Text())
	}

	return scanner.Err()
}

func run() error {
	args := flag.Args()
	// 引数チェック
	if len(args) < 1 {
		return fmt.Errorf("ファイルを指定してください。")
	}

	return cat(args[0])

}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

}
