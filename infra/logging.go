package infra

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

func MainLogging() {
	log.Println("standart logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	fmt.Println("with micro")

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)

}
