# Wails v3 + Nuxt 4 Desktop Template

这是一个基于 **Wails v3** 和 **Nuxt 4** 构建的现代桌面应用模板。它集成了 **Nuxt UI** 和 **Tailwind CSS**，旨在提供极致的开发体验和精美的界面设计。

## 🚀 技术栈

- **后端**: [Wails v3](https://v3.wails.io/) (Go) - 轻量级、高性能的跨平台桌面应用框架。
- **前端**: [Nuxt 4](https://nuxt.com/) (Vue 3) - 强大的全栈 Vue 框架。
- **UI 组件库**: [Nuxt UI](https://ui.nuxt.com/) - 基于 Tailwind CSS 的高质量组件库。
- **样式**: [Tailwind CSS](https://tailwindcss.com/) - 原子化 CSS 框架。
- **包管理器**: [pnpm](https://pnpm.io/) - 快速、节省磁盘空间的包管理器。

## 🛠️ 快速开始

### 前置要求

- [Go](https://go.dev/dl/) (最新版本)
- [Node.js](https://nodejs.org/) (建议 v18+)
- [pnpm](https://pnpm.io/installation)
- [Wails v3 CLI](https://v3.wails.io/getting-started/installation/)

### 安装依赖

进入 `frontend` 目录并安装前端依赖：

```bash
cd frontend
pnpm install
```

### 开发模式

在项目根目录下运行以下命令，启动开发服务器（支持前后端热重载）：

```bash
wails3 dev
```

### 构建应用

构建生产环境的可执行文件：

```bash
wails3 build
```

## 📂 项目结构

- `main.go`: 应用入口，配置 Wails 应用实例和绑定的服务。
- `internal/`:
  - `services/`: 存放 Go 业务逻辑服务（可被前端调用）。
  - `windows/`: 窗口配置和管理。
- `frontend/`: Nuxt 前端应用。
  - `app/components/`: Vue 组件。
  - `app/pages/`: 路由页面。
  - `assets/`: 静态资源（图片、样式等）。
  - `nuxt.config.ts`: Nuxt 配置文件。
- `wails.json`: Wails 项目配置文件。

## 💡 开发提示

- **服务绑定**: 在 `internal/services` 中定义的 Go 结构体方法可以通过 `wails3 dev` 自动生成的 JS bindings 在前端调用。
- **UI 组件**: 本模板已内置 Nuxt UI，你可以直接使用 `<UButton>`, `<UInput>`, `<UCard>` 等组件。
- **图标**: 支持使用 Iconify 图标，例如 `<UIcon name="i-lucide-home" />`。
