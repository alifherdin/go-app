package domains

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `gorm:"type:uuid; default:gen_random_uuid(); not null"`
	Email        string    `gorm:"unique; not null; default:null"`
	Username     string    `gorm:"unique; default:null"`
	PasswordHash string    `gorm:"type:varchar(255); not null; default:null"`
}
