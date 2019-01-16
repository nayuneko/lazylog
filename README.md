# Lazy logger

## 究極に面倒くさい人のためのロガー

- fmt.Print*ぐらいの温度感で使いたいけどログ日時は出力したい
- エラーログは標準エラーで出力したい
- ログレベルで出し分けしたい
- fmt.Printfで改行コードを入れるのが面倒くさい

# Usage

    import "github.com/nayuneko/lazylog"

    func main() {

      // 直接呼び出す
      lazylog.Info("logging start...")            // "YYYY/MM/DD HH:ii:ss I logging start..."
      lazylog.Infof("logging end: count: %d", 10) // "YYYY/MM/DD HH:ii:ss I logging end: count: 10"

      // ログレベルで制御する
      logger := lazylog.NewInfo() // lazylog.New({loglevel: lazylog.loglevelInfo})でもOK!
      logger.Error("errorlog") // "YYYY/MM/DD HH:ii:ss E errorlog"
      logger.Info("infolog")   // "YYYY/MM/DD HH:ii:ss I infolog"
      logger.Debug("debuglog") // 出力なし

    }

# loglevel

ログレベルは5つ！

| ログレベル | 概要 | ログ出力先 | 呼び出し後の動作 |
|:-:|:-:|:-:|:-:|
|loglevelFatal|致命的エラー|os.Stderr|os.Exit(1)を呼び出し終了|
|loglevelError|エラー|os.Stderr|特になし|
|loglevelInfo|情報|os.Stdout|特になし|
|loglevelDebug|デバッグ情報|os.Stdout|特になし|
|loglevelTrace|トレース情報|os.Stdout|特になし|
