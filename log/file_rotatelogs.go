package log

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/stewie1520/blog_ent/config"
	"go.uber.org/zap/zapcore"
)

func getFileRotateWriteSyncer() (zapcore.WriteSyncer, error) {
	filewriter, err := rotatelogs.New(
		path.Join("./tmp", "%Y-%m-%d.log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(7*24*time.Hour)),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if !config.C().IsProd() {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(filewriter)), err
	}

	return zapcore.AddSync(filewriter), err
}
