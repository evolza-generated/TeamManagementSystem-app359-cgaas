package dto

type Roles struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    RolesId string `json:"RolesId" validate:"required"`
    Name string `json:"Name" validate:"required"`
    Description string `json:"Description" validate:"required"`
    CreatedAt string `json:"CreatedAt" validate:"required"`
    UpdatedAt string `json:"UpdatedAt" validate:"required"`
    }

