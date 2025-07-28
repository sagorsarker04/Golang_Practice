package main

import (
    "log"

    "go.uber.org/zap"
)

func main() {
    logger, err := zap.NewProduction()
    if err != nil {
        log.Fatal(err)
    }

    sugar := logger.Sugar()
    defer logger.Sync()

    sugar.Debug("this is a debug message")
    sugar.Info("this is an info message")
    sugar.Warn("this is a warn message")
    sugar.Error("this is an error message")
    sugar.Fatal("this is a fatal message")
}
