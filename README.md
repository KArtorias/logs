# 日志系统使用方法
地址：
-------------
`go get github.com/AgeOfLegends/logs`

功能
----

1.日志分等级：info/warn/error/fatal

2.可传入context，并设置log_key,传入logid从而串联上下文

3.日志可设置自动划分，按时间长短可以设置为天、时、分三个级别

4.可自定义日志文件位置

demo
----
    logs.SetLogSplitType(HOUR)
    logs.SetDir("output/logs/")

	var logKey = "LOG-ID"
	logs.SetLogKey(logKey)

	logs.StartLog()

	ctx := context.Background()
	ctx = context.WithValue(ctx, logKey, "1234567890abcdefghigklmn")
	logs.CtxInfo(ctx, "test info:%d, %s", 1, "info detail")
	logs.CtxWarn(ctx, "test warn:%d, %s", 2, "warn detail")
	logs.CtxError(ctx, "test error:%d, %s", 3, "error detail")
	logs.CtxFatal(ctx, "test fatal:%d, %s", 4, "fatal detail")

	logs.Stop()

运行结果
------
````
2020/10/02 14:36:39 logs.go:53: FATAL D:/Go_work/src/github.com/AgeOfLegends/logs/logs_test.go:22 1234567890abcdefghigklmn test fatal:4, fatal detail
2020/10/02 14:36:39 logs.go:29: WARN D:/Go_work/src/github.com/AgeOfLegends/logs/logs_test.go:20 1234567890abcdefghigklmn test warn:2, warn detail
2020/10/02 14:36:39 logs.go:41: ERROR D:/Go_work/src/github.com/AgeOfLegends/logs/logs_test.go:21 1234567890abcdefghigklmn test error:3, error detail
2020/10/02 14:36:39 logs.go:17: INFO D:/Go_work/src/github.com/AgeOfLegends/logs/logs_test.go:19 1234567890abcdefghigklmn test info:1, info detail
