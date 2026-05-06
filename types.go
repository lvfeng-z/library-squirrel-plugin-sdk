package pluginsdk

// SlotType 插槽类型
type SlotType string

const (
	SlotTypeEmbed           SlotType = "embed"
	SlotTypePanel           SlotType = "panel"
	SlotTypeView            SlotType = "view"
	SlotTypeMenu            SlotType = "menu"
	SlotTypeSiteBrowserList SlotType = "siteBrowserList"
)

// ContentType 插槽内容类型
type ContentType string

const (
	ContentTypeVueSource ContentType = "vueSource"
	ContentTypeHTML      ContentType = "html"
	ContentTypeComponent ContentType = "component"
)

// WindowHandle 窗口句柄
type WindowHandle interface {
	Show()
	Close()
	SetTitle(title string)
}

// WindowOptions 创建窗口的选项
type WindowOptions struct {
	Title  string
	Width  int
	Height int
	URL    string
}

// TaskCreateResult 任务创建结果
type TaskCreateResult struct {
	Succeed       bool
	AddedQuantity int
	Msg           string
}
