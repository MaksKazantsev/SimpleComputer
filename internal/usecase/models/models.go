package models

import "time"

type (
	Mouse struct {
		ID         string    `gorm:"id,primaryKey" json:"id"`
		RoomID     string    `gorm:"room_id" json:"room_id"`
		BuildingID string    `gorm:"building_id" json:"building_id"`
		CreatedAt  time.Time `gorm:"created_at" json:"created_at"`
		UpdatedAt  time.Time `gorm:"updated_at" json:"updated_at"`
		DPI        string    `gorm:"dpi" json:"dpi"`
		Material   string    `gorm:"material" json:"material"`
		Size       string    `gorm:"size" json:"size"`
		Sensor     string    `gorm:"sensor" json:"sensor"`
	}

	GetMouseReq struct {
		ID         string `gorm:"id" json:"id"`
		RoomID     string `gorm:"room_id" json:"room_id"`
		BuildingID string `gorm:"building_id" json:"building_id"`
		DPI        string `gorm:"dpi" json:"dpi"`
		Material   string `gorm:"material" json:"material"`
		Size       string `gorm:"size" json:"size"`
		Sensor     string `gorm:"sensor" json:"sensor"`
	}

	KeyBoard struct {
		ID         string    `gorm:"id,primaryKey" json:"id"`
		RoomID     string    `gorm:"room_id" json:"room_id"`
		BuildingID string    `gorm:"building_id" json:"building_id"`
		CreatedAt  time.Time `gorm:"created_at" json:"created_at"`
		UpdatedAt  time.Time `gorm:"updated_at" json:"updated_at"`
		Material   string    `gorm:"material" json:"material"`
		Size       string    `json:"size"`
		KeyType    string    `json:"key_type"`
		Switches   string    `gorm:"switches" json:"switches"`
	}

	GetKeyBoardReq struct {
		ID         string `gorm:"id" json:"id"`
		RoomID     string `gorm:"room_id" json:"room_id"`
		BuildingID string `gorm:"building_id" json:"building_id"`
		Material   string `gorm:"material" json:"material"`
		Size       string `gorm:"size" json:"size"`
		KeyType    string `json:"key_type"`
		Switches   string `gorm:"switches" json:"switches"`
	}

	Screen struct {
		ID         string    `gorm:"id,primaryKey" json:"id"`
		RoomID     string    `gorm:"room_id" json:"room_id"`
		BuildingID string    `gorm:"building_id" json:"building_id"`
		CreatedAt  time.Time `gorm:"created_at" json:"created_at"`
		UpdatedAt  time.Time `gorm:"updated_at" json:"updated_at"`
		Size       string    `gorm:"size" json:"size"`
		Herz       string    `gorm:"herz" json:"herz"`
		Matrix     string    `gorm:"matrix" json:"matrix"`
	}

	GetScreenReq struct {
		ID         string `gorm:"id" json:"id"`
		RoomID     string `gorm:"room_id" json:"room_id"`
		BuildingID string `gorm:"building_id" json:"building_id"`
		Size       string `gorm:"size" json:"size"`
		Herz       string `gorm:"herz" json:"herz"`
		Matrix     string `gorm:"matrix" json:"matrix"`
	}

	Case struct {
		ID            string    `gorm:"id,primaryKey" json:"id"`
		RoomID        string    `gorm:"room_id" json:"room_id"`
		BuildingID    string    `gorm:"building_id" json:"building_id"`
		CreatedAt     time.Time `gorm:"created_at" json:"created_at"`
		UpdatedAt     time.Time `gorm:"updated_at" json:"updated_at"`
		Size          string    `gorm:"size" json:"size"`
		Color         string    `gorm:"color" json:"color"`
		Compatibility string    `gorm:"compatibility" json:"compatibility"`
	}

	GetCaseReq struct {
		ID            string `gorm:"id" json:"id"`
		RoomID        string `gorm:"room_id" json:"room_id"`
		BuildingID    string `gorm:"building_id" json:"building_id"`
		Size          string `gorm:"size" json:"size"`
		Color         string `gorm:"color" json:"color"`
		Compatibility string `gorm:"compatibility" json:"compatibility"`
	}
)
