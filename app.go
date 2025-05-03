package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) StartDownload() {
	for progress := 0; progress <= 100; progress += 10 {
		runtime.EventsEmit(a.ctx, "download_progress", progress)
		time.Sleep(time.Second)
	}
	runtime.EventsEmit(a.ctx, "download_success")
}
