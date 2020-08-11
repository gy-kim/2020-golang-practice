package cron

type Logger interface {
	Info(msg string, keysAndValues ...interface{})

	Error(err error, msg string, keysAndValues ...interface{})
}

// func PrintfLogger(l interface{ Printf(string, ...interface{}) }) Logger {
// 	return printfLogger{l, false}
// }

// func VerbosePrintfLogger(l interface{ Printf(string, ...interface{}) }) Logger {
// 	return printfLogger{l, true}
// }

// type printfLogger struct {
// 	logger  interface{ Printf(string, ...interface{}) }
// 	logInfo bool
// }

// func (pl printfLogger) Info(msg string, keysAndValues ...interface{}) {
// 	if pl.logInfo {
// 		keysAndValues = formatTimes(keysAndValues)

// 	}
// }
