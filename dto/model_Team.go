package dto

type Team struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    TeamId string `json:"TeamId" validate:"required"`
    Name string `json:"Name" validate:"required"`
    Description string `json:"Description" validate:"required"`
    CreatedAt string `json:"CreatedAt" validate:"required"`
    UpdatedAt string `json:"UpdatedAt" validate:"required"`
    }

