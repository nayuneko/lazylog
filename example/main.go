package main

import "github.com/nayuneko/lazylog"

func main() {

	// 直接呼び出す
	lazylog.Info("logging start...") // "YYYY/MM/DD HH:ii:ss I logging start..."
	lazylog.Infof("count: %d", 10)   // "YYYY/MM/DD HH:ii:ss I count: 10"
	lazylog.Trace("traceinfo")       // "YYYY/MM/DD HH:ii:ss T traceinfo"
	lazylog.Debug("debuginfo")       // "YYYY/MM/DD HH:ii:ss D debuginfo"
	lazylog.Error("errorinfo")       // "YYYY/MM/DD HH:ii:ss E errorinfo"
	//lazylog.Fatal("fatal")           // "YYYY/MM/DD HH:ii:ss E fatal"

	// ログレベルで制御する
	logger := lazylog.NewInfo() // lazylog.New({loglevel: lazylog.LoglevelInfo})でもOK!
	logger.Error("errorlog")    // "YYYY/MM/DD HH:ii:ss E errorlog"
	logger.Info("infolog")      // "YYYY/MM/DD HH:ii:ss I infolog"
	logger.Debug("debuglog")    // 出力なし
	logger.Trace("tracelog")    // 出力なし

}
