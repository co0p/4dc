package tray

type Item struct {
	Title   string
	OnClick func()
}

type Tray interface {
	Show(title string, items []Item)
}
