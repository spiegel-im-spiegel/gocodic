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

## パッケージ機能

[gocodic] では以下の機能を提供します。

- ネーミング翻訳機能 ([transration.md](transration.md))
- プロジェクト参照機能 ([projects.md](projects.md))
- CED ルックアップ機能 ([lookup.md](lookup.md))

詳しくは括弧内のドキュメントを参照してください。

## その他

日本人なのでフィードバック等は日本語でおｋ。
中身についてはブログでちょっとだけ解説しています。

- [Codic API を利用するパッケージを作ってみた — プログラミング言語 Go | text.Baldanders.info](http://text.baldanders.info/golang/codic-api/)

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
