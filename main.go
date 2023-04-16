package main

import (
	"github.com/alexflint/go-arg"
	"log"
	"problem-1/gen"
	"problem-1/lib"
	"strconv"
	"time"
)

var args struct {
	Listen *ListenCmd `arg:"subcommand:listen"`
	Gen    *GenCmd    `arg:"subcommand:generate"`
}

type ListenCmd struct {
	Port int    `default:"1321"`
	File string `arg:"required"`
}

type GenCmd struct {
	Num  int    `arg:"required"`
	File string `arg:"required"`
}

func main() {

	arg.MustParse(&args)

	switch {
	case args.Gen != nil:
		gen.GeneratePromotionsCsvFile(args.Gen.File, args.Gen.Num)
	case args.Listen != nil:
		listen()
	}

}

func listen() {
	go lib.InitStorage(args.Listen.File+".csv", 30*time.Minute)
	lib.RunHttpServer(strconv.Itoa(args.Listen.Port))

	defer func() {
		log.Print("Closing application")
		lib.CloseStorage()
	}()
}
