# プロジェクト参照機能 | [gocodic] パッケージ

＜ [README.md](README.md) に戻る

プロジェクト参照機能では [gocodic] においてユーザが所属（オーナーも含む）するプロジェクトの情報を取得します。

## パッケージの利用例

### プロジェクト一覧の取得

まず，以下に示すように， `opts *options.Options` を生成して必要な情報をセットします。

```go
opts, err := options.NewOptions(options.CmdProjLst, "YOUR_ACCESS_TOKEN")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

`YOUR_ACCESS_TOKEN` の部分には [codic] サイトで取得したアクセストークンをセットします。
実際に API に対してリクエストを送る処理は以下のとおりです。

```go
res, err := gocodic.ReferProjects(opts, 0)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

`gocodic.ReferProjects()` 関数の第2引数には任意の値でかまいません。

返り値の `res *response.Response` には API からのレスポンスが格納されます。
リクエストが成功した場合と失敗した場合で応答データのフォーマットが異なため，各々で場合分けして処理します。

```go
if res.IsSuccess() {
    sd, err := response.DecodeSuccessProjects(res.Body())
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    for _, d := range sd {
        fmt.Printf("%d:%s, %s (Owner: %d:%s)\n", d.ID, d.Name, d.Description, d.Owner.ID, d.Owner.Name)
    }
} else {
    ed, err := response.DecodeError(res.Body())
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    for _, d := range ed.Errors {
        fmt.Fprintln(os.Stderr, d.Message)
    }
}
```

これで，リクエストが成功した場合は `sd []response.SuccessProject` に，失敗した場合は `ed *response.ErrorList` に結果が格納されます。

### プロジェクト詳細の取得

まず，以下に示すように， `opts *options.Options` を生成して必要な情報をセットします。

```go
opts, err := options.NewOptions(options.CmdProj, "YOUR_ACCESS_TOKEN")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

`YOUR_ACCESS_TOKEN` の部分には [codic] サイトで取得したアクセストークンをセットします。
実際に API に対してリクエストを送る処理は以下のとおりです。

```go
res, err := gocodic.ReferProjects(opts, 123)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

`gocodic.ReferProjects()` 関数の第2引数には参照するプロジェクトの ID 番号をセットします。

返り値の `res *response.Response` には API からのレスポンスが格納されます。
リクエストが成功した場合と失敗した場合で応答データのフォーマットが異なため，各々で場合分けして処理します。

```go
if res.IsSuccess() {
    sd, err := response.DecodeSuccessProjects(res.Body())
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
	fmt.Printf("%d:%s, %s (Owner: %d:%s)\n", sd.ID, sd.Name, sd.Description, sd.Owner.ID, sd.Owner.Name)
} else {
    ed, err := response.DecodeError(res.Body())
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    for _, d := range ed.Errors {
        fmt.Fprintln(os.Stderr, d.Message)
    }
}
```

これで，リクエストが成功した場合は `sd *response.SuccessProject` に，失敗した場合は `ed *response.ErrorList` に結果が格納されます。

## [gocodic] コマンドライン・インタフェース： proj サブコマンド

[gocodic] ではコマンドライン・インタフェースも用意しています。

```
$ gocodic
APIs for codic.jp

Usage:
  gocodic [command]

Available Commands:
  help        Help about any command
  lookup      Lookup CED API for codic.jp
  proj        Refer projects API for codic.jp
  trans       Transration API for codic.jp
  version     Print the version number of gocodic

Flags:
      --config string   config file (default is $HOME/.gocodic.yaml)
  -h, --help            help for gocodic
  -j, --json            output by JSON format (raw data)
  -t, --token string    access token of codic.jp

Use "gocodic [command] --help" for more information about a command.
```

このうちプロジェクト参照機能を提供するのが `proj` サブコマンドです。

```
$ gocodic proj -h
Refer projects API for codic.jp

Usage:
  gocodic proj [flags]

Flags:
  -h, --help            help for proj
  -p, --projectid int   project_id parameter

Global Flags:
      --config string   config file (default is $HOME/.gocodic.yaml)
  -j, --json            output by JSON format (raw data)
  -t, --token string    access token of codic.jp
```

プロジェクト一覧を取得する場合は以下のように使います。

```
$ gocodic proj -t xxxx
8099:My Project, お試し用プロジェクト (Owner: 9131:Der Spiegel im Spiegel)
```

各プロジェクトの詳細が見たい場合はプロジェクト ID を指定して

```
$ gocodic proj -t xxxx -p 8099
8099:My Project, お試し用プロジェクト (Owner: 9131:Der Spiegel im Spiegel)
```

とします。
`-j` オプションで JSON 形式のまま出力することも出来ます。

```
$ gocodic proj -t xxxx -p 8099 -j
{"id":8099,"name":"My Project","description":"\u304a\u8a66\u3057\u7528\u30d7\u30ed\u30b8\u30a7\u30af\u30c8","words_count":0,"created_on":"Tue, 24 Oct 2017 17:37:03 +0900","owner":{"id":9131,"name":"Der Spiegel im Spiegel"}}
```

アクセストークンはホームディレクトリ直下の設定ファイル `.gocodic.yaml` に記述しておけば起動時に設定を読み込みます。
設定ファイルの記述例は以下のとおりです（YAML フォーマット）。

```yaml
token: YOUR_ACCESS_TOKEN
casing: camel
style: camel strict
projectid: 123
```

設定ファイルは `--config` フラグで指定することも出来ます。
なお，設定ファイルの内容よりもコマンドラインの引数のほうが優先されます。
また `proj` サブコマンドでは `projectid` フラグはコマンドラインの値のみ有効で設定ファイルの値は無視されます。

＜ [README.md](README.md) に戻る

[gocodic]: https://github.com/spiegel-im-spiegel/gocodic "spiegel-im-spiegel/gocodic: codic の API を利用するための Go 言語パッケージ"
[codic]: https://codic.jp/ "プログラマーのためのネーミング辞書 | codic"
