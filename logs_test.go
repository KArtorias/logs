package logs

import (
	"context"
	"testing"
)

func TestLogs(t *testing.T) {
	SetLogSplitType(HOUR)
	SetDir("output/logs/")

	var logKey = "LOG-ID"
	SetLogKey(logKey)

	StartLog()

	ctx := context.Background()
	ctx = context.WithValue(ctx, logKey, "1234567890abcdefghigklmn")
	CtxInfo(nil, "test info:%d, %s", 1, "info detail")
	CtxWarn(ctx, "test warn:%d, %s", 2, "warn detail")
	CtxError(ctx, "test error:%d, %s", 3, "error detail")
	CtxFatal(ctx, "test fatal:%d, %s", 4, "fatal detail")

	Stop()
}
