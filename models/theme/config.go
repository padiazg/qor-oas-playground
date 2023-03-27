package theme

type Config struct {
	ID      string `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"unique"`
	Author  string `json:"author"`
	Version string `json:"version"`
	Path    string `json:"path"`
	Status  bool   `json:"status" gorm:"default:false"`
	// ThemeFile File   `gorm:"-"`
}
