// This code is part of player package
package orgModel

import (
	"time"

	"gorm.io/gorm"
)

// import other packages

type Org struct {
	OrganisationOwnerID string `json:"OrganisationOwnerID"`
	Name                string `json:"Name"`
	OrganisationID      string `gorm:"uniqueIndex"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}
