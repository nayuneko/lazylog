# Lazy logger

### 究極に面倒くさい人のためのGo言語用ロガー

- `fmt.Println`,`fmt.Printf`ぐらいの温度感でロガーを使いたい
- ログ日時は出力したい
- エラーログは標準エラーで出力したい
- ログレベルで出し分けしたい
- fmt.Printfで改行コードを入れるのが面倒くさい

# Installation

    go get github.com/nayuneko/lazylog

# Usage

    import "github.com/nayuneko/lazylog"

    func main() {

      // 直接呼び出す
      lazylog.Info("logging start...")            // "YYYY/MM/DD HH:ii:ss I logging start..."
      lazylog.Infof("logging end: count: %d", 10) // "YYYY/MM/DD HH:ii:ss I logging end: count: 10"

      // ログレベルで制御する
      logger := lazylog.NewInfo() // alias: lazylog.New({loglevel: lazylog.LoglevelInfo})
      logger.Error("errorlog") // "YYYY/MM/DD HH:ii:ss E errorlog"
      logger.Info("infolog")   // "YYYY/MM/DD HH:ii:ss I infolog"
      logger.Debug("debuglog") // 出力なし

      loggerFatal := lazylog.NewFatal() // alias: lazylog.New({loglevel: lazylog.LoglevelFatal})
      logger.Error("errorlog") // 出力なし
      logger.Info("infolog")   // 出力なし
      logger.Debug("debuglog") // 出力なし

    }

# loglevel

ログレベルは5つ！

| ログレベル | 概要 | ログ出力先 | 呼び出し後の動作 |
|:-:|:-:|:-:|:-:|
|Fatal|致命的エラー|os.Stderr|os.Exit(1)を呼び出し終了|
|Error|エラー|os.Stderr|特になし|
|Info|情報|os.Stdout|特になし|
|Debug|デバッグ情報|os.Stdout|特になし|
|Trace|トレース情報|os.Stdout|特になし|

NewXXXでロガーのインスタンスを生成した場合は指定したログレベル以上のログのみ出力されます。