package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	numberFlag bool
)

func init() {
	flag.BoolVar(&numberFlag, "n", false, "show row numbre")
	flag.Parse()
}

func cat(file ...string) error {
	for j := 0; j < len(file); j++ {

		fi, err := os.Open(file[j])
		if err != nil {
			return fmt.Errorf("ファイルをひらけませんでした。")
		}
		defer fi.Close()

		scanner := bufio.NewScanner(fi)
		if numberFlag {
			// sfから1行ずつ読み込み、"行数:"を前に付けて標準出力に書き出す。
			for i := 1; scanner.Scan(); i++ {
				fmt.Fprintf(os.Stdout, "%d : %s\n", i, scanner.Text())
			}
		} else {
			for i := 1; scanner.Scan(); i++ {
				fmt.Fprintf(os.Stdout, "%s\n", scanner.Text())
			}
		}
	}
	return nil
}

func run() error {
	args := flag.Args()
	// 引数チェック
	if len(args) < 1 {
		return fmt.Errorf("ファイルを指定してください。")
	}

	return cat(args...)

}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

}
