# CED ルックアップ機能 | [gocodic] パッケージ

＜ [README.md](README.md) に戻る

CED ルックアップ機能では [gocodic] 辞書である CED (Codic English Dictionary) の情報を取得します。

## パッケージの利用例

## CED クエリ発行

まず，以下に示すように， `opts *options.Options` を生成して必要な情報をセットします。

```go
opts, err := options.NewOptions(options.CmdCEDQuery, "YOUR_ACCESS_TOKEN")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
opts.Add(options.Query("term"))
opts.Add(options.Count(2))
```

`YOUR_ACCESS_TOKEN` の部分には [codic] サイトで取得したアクセストークンをセットします。
実際に API に対してリクエストを送る処理は以下のとおりです。

```go
res, err := gocodic.LookupCED(opts, 0)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

`gocodic.LookupCED()` 関数の第2引数には任意の値でかまいません。

返り値の `res *response.Response` には API からのレスポンスが格納されます。
リクエストが成功した場合と失敗した場合で応答データのフォーマットが異なため，各々で場合分けして処理します。

```go
if res.IsSuccess() {
    sd, err := response.DecodeSuccessLookup(res.Body())
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    for _, d := range sd {
        fmt.Printf("%d:%s, %s\n", d.ID, d.Title, d.Digest)
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

これで，リクエストが成功した場合は `sd []response.SuccessLookup` に，失敗した場合は `ed *response.ErrorList` に結果が格納されます。

## CED エントリの取得

まず，以下に示すように， `opts *options.Options` を生成して必要な情報をセットします。

```go
opts, err := options.NewOptions(options.CmdCED, "YOUR_ACCESS_TOKEN")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

`YOUR_ACCESS_TOKEN` の部分には [codic] サイトで取得したアクセストークンをセットします。
実際に API に対してリクエストを送る処理は以下のとおりです。

```go
res, err := gocodic.LookupCED(opts, 123)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

`gocodic.LookupCED()` 関数の第2引数には参照する CED エントリの ID 番号をセットします。

返り値の `res *response.Response` には API からのレスポンスが格納されます。
リクエストが成功した場合と失敗した場合で応答データのフォーマットが異なため，各々で場合分けして処理します。

```go
if res.IsSuccess() {
    sd, err := response.DecodeSuccessProjects(res.Body())
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Printf("%d:%s, %s\n", sd.ID, sd.Title, sd.Digest)
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

これで，リクエストが成功した場合は `sd *response.SuccessEntry` に，失敗した場合は `ed *response.ErrorList` に結果が格納されます。

## [gocodic] コマンドライン・インタフェース： lookup サブコマンド

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

このうちCED ルックアップ機能を提供するのが `lookup` サブコマンドです。

```
$ gocodic lookup -h
Lookup CED API for codic.jp

Usage:
  gocodic lookup [flags] <query string>

Flags:
  -c, --count int     count parameter
  -e, --entryid int   CED entry ID
  -h, --help          help for lookup

Global Flags:
      --config string   config file (default is $HOME/.gocodic.yaml)
  -j, --json            output by JSON format (raw data)
  -t, --token string    access token of codic.jp
```

CED に対してクエリを発行する場合は以下のように使います。

```
$ gocodic lookup -t xxxx -c 2 term
41941:term, 用語、期間、期限
43875:terminal, 末端の、末端、端末
```

各エントリの詳細が見たい場合はエントリ ID を指定して

```
$ gocodic lookup -t xxxx -e 43875
43875:terminal, 末端の、末端、端末
```

とします。
`-j` オプションで JSON 形式のまま出力することも出来ます。

```
$ gocodic lookup -t xxxx -e 43875 -j
{"id":43875,"title":"terminal","digest":"\u672b\u7aef\u306e\u3001\u672b\u7aef\u3001\u7aef\u672b","pronunciations":[{"type":"katakana","text":"\u30bf\u30fc\u30df\u30ca\u30eb","labels":[]}],"translations":[{"note":null,"labels":[],"etymology":1,"pos":"noun","text":"\uff08\u7a7a\u6e2f\u306e\uff09\u30bf\u30fc\u30df\u30ca\u30eb"},{"note":null,"labels":[],"etymology":1,"pos":"noun","text":"\uff08\u96fb\u8eca\u306e\uff09\u30bf\u30fc\u30df\u30ca\u30eb\u99c5"},{"note":null,"labels":[],"etymology":1,"pos":"noun","text":"\u672b\u7aef"},{"note":null,"labels":["computing"],"etymology":1,"pos":"noun","text":"\u7aef\u672b"},{"note":null,"labels":["computing"],"etymology":1,"pos":"noun","text":"\u30bf\u30fc\u30df\u30ca\u30eb\u30a8\u30df\u30e5\u30ec\u30fc\u30bf"},{"note":null,"labels":[],"etymology":1,"pos":"adjective","text":"\u672b\u7aef\u306e"}],"note":""}
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

＜ [README.md](README.md) に戻る

[gocodic]: https://github.com/spiegel-im-spiegel/gocodic "spiegel-im-spiegel/gocodic: codic の API を利用するための Go 言語パッケージ"
[codic]: https://codic.jp/ "プログラマーのためのネーミング辞書 | codic"
