
package logger

import (
    "io"
    "log"
    "os"

    "github.com/gin-gonic/gin"
)

type Logger struct {
    *log.Logger
}

func New(env string) *Logger {
    var out io.Writer = os.Stdout
    l := log.New(out, "", log.LstdFlags|log.Lshortfile)
    if env == "production" {
        gin.SetMode(gin.ReleaseMode)
    }
    return &Logger{l}
}

func GinLogger(l *Logger) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        l.Printf("%s %s %d", c.Request.Method, c.FullPath(), c.Writer.Status())
    }
}
