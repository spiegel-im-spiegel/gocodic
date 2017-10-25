# [gocodic] - [codic] の API を利用するための Go 言語パッケージ

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/spiegel-im-spiegel/gocodic/blob/master/LICENSE)


[gocodic] は [codic] で提供される API を [Go 言語]で利用するためのパッケージです。
API の機能については以下を参照してください。

- [API | codic](https://codic.jp/docs/api)

## 導入

`go get` コマンドで `GOPATH` 配下に [gocodic] パッケージが展開されます。

```
$ go get -v github.com/spiegel-im-spiegel/gocodic
```

`GOPATH` に組み込むのではなく [dep] の制御下に入れる場合は以下のようにします。

```
$ dep ensure -add github.com/spiegel-im-spiegel/gocodic
```

## パッケージの利用例

まず，以下に示すように， `opts *options.Options` を生成して必要な情報をセットします。

```go
opts, err := options.NewOptions("YOUR_ACCESS_TOKEN")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
casing, err := options.NewCasingOption("camel")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
opts.Add(casing)
style, err := options.NewAcronymStyleOption("camel strict")
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
opts.Add(style)
opts.Add(options.Text("ユーザを登録する"))
```

`YOUR_ACCESS_TOKEN` の部分には [codic] サイトで取得したアクセス・トークンをセットします。

`options.NewCasingOption()` の引数には `casing` パラメータの値をセットします。
現時点では以下の値のみ受け付けます。

- `"camel"` (`CaseCamel`)
- `"pascal"` (`CasePascal`)
- `"underscore"` (`CaseUnderscore`)
- `"upper underscore"` (`CaseUpperUnderscore`)
- `"hyphen"` (`CaseHyphen`)

括弧内のシンボルはそれぞれの値を示す定数で `opts.Add(CaseCamel)` のように直接セットできます。

`options.NewAcronymStyleOption()` の引数には `acronym_style` パラメータの値をセットします。
現時点では以下の値のみ受け付けます。

- `"MS naming"` (`StyleMSNaming`)
- `"guidelines"` (`StyleGuidelines`)
- `"camel strict"` (`StyleCamelStrict`)
- `"literal"` (`StyleLiteral`)

括弧内のシンボルはそれぞれの値を示す定数で `opts.Add(StyleCamelStrict)` のように直接セットできます。

実際に API に対してリクエストを送る処理は以下のとおりです。

```go
res, err := gocodic.Translate(opts)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

返り値の `res *response.Response` には API からのレスポンスが格納されます。
リクエストが成功した場合と失敗した場合で応答データのフォーマットが異なため，各々で場合分けして処理します。

```go
if res.IsSuccess() {
    sd, err := response.DecodeSuccessTrans(res.Body())
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    for _, d := range sd {
        fmt.Println(d.TranslatedText)
        //output:
        //registerUser
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

これで，リクエストが成功した場合は `sd []SuccessTrans` に，失敗した場合は `ed *ErrorList` に結果が格納されます。

## コマンドライン・インタフェース

[gocodic] ではコマンドライン・インタフェースも用意しています。

```
$ gocodic
APIs for codic.jp

Usage:
  gocodic [command]

Available Commands:
  help        Help about any command
  trans       Ttansration API for codic.jp

Flags:
      --config string   config file (default is $HOME/.gocodic.yaml)
  -h, --help            help for gocodic
  -j, --json            output by JSON format (raw data)
  -t, --token string    access token of codic.jp

Use "gocodic [command] --help" for more information about a command.

$ gocodic trans -h
Ttansration API for codic.jp

Usage:
  gocodic trans [flags] [<word>...]

Flags:
  -c, --casing string   casing parameter
  -h, --help            help for trans
  -p, --projectid int   project_id parameter
  -s, --style string    acronym_style parameter

Global Flags:
      --config string   config file (default is $HOME/.gocodic.yaml)
  -j, --json            output by JSON format (raw data)
  -t, --token string    access token of codic.jp
```

たとえば以下のように使います。

```
$ gocodic trans -t xxxx -c camel -s "camel strict" ユーザを登録する 登録したか
registerUser
isRegistered
```

`-j` オプションで JSON 形式のまま出力することも出来ます。

```
$ gocodic trans -t xxxx -c camel -s "camel strict" ユーザを登録する
[{"successful":true,"text":"\u30e6\u30fc\u30b6\u3092\u767b\u9332\u3059\u308b","translated_text":"registerUser","words":[{"successful":true,"text":"\u767b\u9332\u3059\u308b","translated_text":"register","candidates":[{"text":"register"},{"text":"registering"},{"text":"join"},{"text":"to register"}]},{"successful":true,"text":"\u30e6\u30fc\u30b6","translated_text":"user","candidates":[{"text":"user"}]},{"successful":true,"text":"\u3092","translated_text":null,"candidates":[{"text":null},{"text":"that"},{"text":"to"},{"text":"for"},{"text":"from"},{"text":"is"},{"text":"of"}]}]}]
```

さらに翻訳対象の言葉をパイプで渡すことも出来ます。

```
$ cat input.txt
ユーザを登録する
登録したか

$ cat input.txt | gocodic trans -t xxxx -c camel -s "camel strict"
registerUser
isRegistered
```


各フラグはホームディレクトリ直下の設定ファイル `.gocodic.yaml` に記述しておけば起動時に設定を読み込みます。
設定ファイルの記述例は以下のとおりです（YAML フォーマット）。

```yaml
token: YOUR_ACCESS_TOKEN
casing: camel
style: camel strict
```

これで先程のコマンドラインをフラグ指定なしで起動できます。

```
$ cat input.txt
ユーザを登録する
登録したか

$ cat input.txt | gocodic trans
registerUser
isRegistered
```

設定ファイルは `--config` フラグで指定することも出来ます。
なお，設定ファイルの内容よりもコマンドラインの引数のほうが優先されます。

## その他

日本人なので日本語でおｋ。

## 参考

- [プログラマーのためのネーミング辞書 | codic](https://codic.jp/)
- [codic-project/Codic_cli](https://github.com/codic-project/Codic_cli) : [Go 言語]による別実装
- [39e/go-codic](https://github.com/39e/go-codic) : [Go 言語]による別実装
- [【codic】プログラマ必見！もう変数名や関数名に困らない！プログラマのためのネーミングツールを紹介 - プログラミング向上雑記](http://niisi.hatenablog.jp/entry/2016/08/17/171000)
- [関数や変数のネーミングに悩んだら「codic」に日本語名を入力するとある程度解決するかも](https://nelog.jp/codic)

[gocodic]: https://github.com/spiegel-im-spiegel/gocodic "spiegel-im-spiegel/gocodic: codic の API を利用するための Go 言語パッケージ"
[codic]: https://codic.jp/ "プログラマーのためのネーミング辞書 | codic"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"
[Go 言語]: https://golang.org/ "The Go Programming Language"
