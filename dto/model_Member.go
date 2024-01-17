package dto

type Member struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    MemberId string `json:"MemberId" validate:"required"`
    Name string `json:"Name" validate:"required"`
    Age int `json:"Age" validate:"required"`
    Email string `json:"Email" validate:"required"`
    Phone string `json:"Phone" validate:"required"`
    Address string `json:"Address" validate:"required"`
    Gender string `json:"Gender" validate:"required"`
    CreatedAt string `json:"CreatedAt" validate:"required"`
    }

