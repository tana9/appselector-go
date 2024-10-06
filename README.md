# AppSelector-go

指定したファイルを開くアプリをFuzzyFinderで選択するツール

## 設定ファイル

config.toml

~~~toml
apps = ["notepad", "goneovim", "code", "zed"]
~~~

## あふwから呼び出す

~~~
AppSelector.go $P\$F
~~~

* `$P` : 時ファイル窓のパス名
* `$F` : カーソル行のファイル名