package entity

import (
	"bytes"
	"context"
)

type FileCID struct {
	ID   uint    `json:"id" gorm:"primaryKey;auto_increment" binding:"-" example:"1"`                            // ID
	Name string  `json:"name" gorm:"type:varchar(255);not null;" example:"meeting.txt"`                          // name
	CID  *string `json:"cid" gorm:"type:varchar(255);" example:"QmWMgCrvNEoD6KnqFfnv4tz7X6soDJvioswNBjHt6XcEum"` // cid of ipfs
}

type FileCIDService interface {
	GetFileCID(ctx context.Context, id uint) (*[]byte, error)
	CreateFileCID(ctx context.Context, buf *bytes.Buffer, name string) (*FileCID, error)
}

type FileCIDRepository interface {
	FindFirst(ctx context.Context, field string, value any) (*FileCID, error)
	Create(ctx context.Context, entity FileCID) (*FileCID, error)
}
