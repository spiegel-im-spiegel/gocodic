# [gocodic] - [codic] の API を利用するための Go 言語パッケージ

[gocodic] は [codic] で提供される API を [Go 言語]で利用するためのパッケージです。

## 導入

以下のコマンドで `GOPATH` 上に [gocodic] パッケージが展開されます。

```
$ go get -v github.com/spiegel-im-spiegel/gocodic
```

[dep] の制御下に入れる場合は以下のようにします。

```
$ dep ensure -add github.com/spiegel-im-spiegel/gocodic
```

## パッケージの利用例

まずは `options.Options` に必要な情報をセットします。

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

- `"camel"`
- `"pascal"`
- `"underscore"`
- `"upper underscore"`
- `"hyphen"`

`options.NewAcronymStyleOption()` の引数には `acronym_style` パラメータの値をセットします。
現時点では以下の値のみ受け付けます。

- `"MS naming"`
- `"guidelines"`
- `"camel strict"`
- `"literal"`

実際に API に対してリクエストを送る処理は以下のとおりです。

```go
res, err := gocodic.Translate(opts)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

返り値の `response.Response` には API からのレスポンスが以下の形式で格納されます。

```go
//Response class is response data from codic service
type Response struct {
    statusCode int
    status     string
    body       []byte
}
```

`body` には JSON 形式のテキストが格納されますが，リクエストが成功した場合と失敗した場合でフォーマットが異なためそれぞれで場合分けして処理します。

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

## コマンドライン・インタフェース

[gocodic] ではコマンドライン・インタフェースも用意しています。

```
$ gocodic
APIs for codic.jp

Usage:
  gocodic [command]

Available Commands:
  help        Help about any command
  trans       Ttansration API  for codic.jp

Flags:
      --config string   config file (default is $HOME/.gocodic.yaml)
  -h, --help            help for gocodic
  -t, --token string    access token of codic.jp

Use "gocodic [command] --help" for more information about a command.

$ gocodic trans -h
Ttansration API for codic.jp

Usage:
  gocodic trans [flags] <word>

Flags:
  -c, --casing string   casing parameter
  -h, --help            help for trans
  -p, --projectid int   project_id parameter
  -s, --style string    acronym_style parameter

Global Flags:
      --config string   config file (default is $HOME/.gocodic.yaml)
  -t, --token string    access token of codic.jp
```

たとえば以下のように使います。

```
$ gocodic trans -t xxxx -c camel -s "camel strict" ユーザを登録する
registerUser
```

各フラグはホームディレクトリ直下の `.gocodic.yaml` ファイルに記述することでコマンドライン指定を省略できます。

`.gocodic.yaml` の記述例は以下のとおりです。

```yaml
token: YOUR_ACCESS_TOKEN
casing: camel
style: camel strict
```

## 参考

- [プログラマーのためのネーミング辞書 | codic](https://codic.jp/)
    - [API | codic](https://codic.jp/docs/api)
    - [codic-project/Codic_cli](https://github.com/codic-project/Codic_cli) : [Go 言語]による別実装
    - [39e/go-codic](https://github.com/39e/go-codic) : [Go 言語]による別実装
- [【codic】プログラマ必見！もう変数名や関数名に困らない！プログラマのためのネーミングツールを紹介 - プログラミング向上雑記](http://niisi.hatenablog.jp/entry/2016/08/17/171000)
- [関数や変数のネーミングに悩んだら「codic」に日本語名を入力するとある程度解決するかも](https://nelog.jp/codic)

[gocodic]: https://github.com/spiegel-im-spiegel/gocodic "spiegel-im-spiegel/gocodic: codic の API を利用するための Go 言語パッケージ"
[codic]: https://codic.jp/ "プログラマーのためのネーミング辞書 | codic"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"
[Go 言語]: https://golang.org/ "The Go Programming Language"
