package windows

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

// MainWindowConfig 主窗口配置
type MainWindowConfig struct {
	Width  int
	Height int
}

// CreateMainWindow 创建主窗口
func CreateMainWindow(app *application.App) *application.WebviewWindow {
	mainWindow := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "主页",
		Name:             "main",
		Width:            1024,
		Height:           768,
		MinWidth:         800,
		MinHeight:        600,
		AlwaysOnTop:      false,
		Frameless:        false,
		Hidden:           false,
		URL:              "/",
		BackgroundType:   application.BackgroundTypeSolid,
		BackgroundColour: application.RGBA{Red: 255, Green: 255, Blue: 255, Alpha: 255},
	})

	return mainWindow
}
