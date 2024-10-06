package main

import (
	"flag"
	"fmt"
	"github.com/koki-develop/go-fzf"
	"os"
	"os/exec"
	"path/filepath"
)

const configFile = "config.toml"

// アプリ選択
func selectApp(apps []string) (string, error) {
	f, err := fzf.New(fzf.WithInputPlaceholder("アプリ選択..."))
	if err != nil {
		return "", err
	}

	idxs, err := f.Find(apps, func(i int) string { return apps[i] })
	if err != nil {
		return "", err
	}
	return apps[idxs[0]], nil
}

func startApp(app string, path string) error {
	c := exec.Command(app, path)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Start()
}

func run(path string) error {
	if path == "" {
		return fmt.Errorf("引数にファイルパスが指定されていません")
	}

	// 設定ファイル読み込み
	config, err := LoadConfig(configFile)
	if err != nil {
		return err
	}

	// アプリ選択
	selectedApp, err := selectApp(config.Apps)
	if err != nil {
		return err
	}

	// アプリ実行
	return startApp(selectedApp, path)
}

func exePath() string {
	exePath, _ := os.Executable()
	return filepath.Dir(exePath)
}

func main() {
	_ = os.Chdir(exePath())
	flag.Parse()
	path := flag.Arg(0)

	if err := run(path); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
