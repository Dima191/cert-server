package main

import (
	"context"
	"flag"
	"github.com/Dima191/cert-server/internal/app"
	"log/slog"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"sync"
	"syscall"
)

var (
	isDebug          bool
	memProfile       bool
	cpuProfile       bool
	goroutineProfile bool
	configPath       string
)

func init() {
	flag.BoolVar(&isDebug, "debug", false, "Enable debug logs")
	flag.StringVar(&configPath, "config-path", "./config/config.env", "path to config file")

	flag.BoolVar(&cpuProfile, "cpu-profile", false, "Enable cpu profile")
	flag.BoolVar(&memProfile, "mem-profile", false, "Enable mem profile")
	flag.BoolVar(&goroutineProfile, "goroutine-profile", false, "Enable goroutine profile")
}

func main() {
	flag.Parse()

	//MEM
	if memProfile {
		defer func() {
			f, _ := os.Create("mem.pb.gz")
			runtime.GC()
			_ = pprof.WriteHeapProfile(f)
			_ = f.Close()
		}()
	}

	//CPU
	if cpuProfile {
		f, _ := os.Create("cpu.pb.gz")
		_ = pprof.StartCPUProfile(f)
		defer func() {
			pprof.StopCPUProfile()
			_ = f.Close()
		}()
	}

	//GOROUTINE
	if goroutineProfile {
		go func() {
			fg, _ := os.Create("goroutine.pb.gz")
			_ = pprof.Lookup("goroutine").WriteTo(fg, 0)
		}()
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	logger := Logger()

	a, err := app.New(ctx, logger, configPath)
	if err != nil {
		logger.Error(err.Error())
		stop()
		os.Exit(1)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		if err = a.Run(); err != nil {
			logger.Error("failed to run application", slog.String("error", err.Error()))
			stop()
		}
	}()

	<-ctx.Done()
	logger.Info("received signal to shut down the application")

	a.Stop()
	wg.Wait()
}

func Logger() *slog.Logger {
	var h slog.Handler

	switch isDebug {
	case true:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	default:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelWarn,
		})
	}

	logger := slog.New(h)

	return logger
}
