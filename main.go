package main

import (
	"changeme/internal/services"
	"changeme/internal/windows"
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	// 创建应用实例
	app := application.New(application.Options{
		Name:        "wails-app",
		Description: "模块仓库",
		Services: []application.Service{
			application.NewService(&services.GreetService{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// 创建主窗口
	windows.CreateMainWindow(app)

	// 监听退出事件
	app.Event.On("application-exit", func(event *application.CustomEvent) {
		app.Quit()
	})

	// 运行应用
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
