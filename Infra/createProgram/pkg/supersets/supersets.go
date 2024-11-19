// This code is part of player package
package supersetsModel

// import other packages

type Superset struct {
	ID         string `gorm:"uniqueIndex"`
	SupersetId string `json:"SupersetId"`
}
