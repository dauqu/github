package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Username   string    `json:"username"`
	FullName   string    `json:"fullname"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Password   string    `json:"password"`
	LicenseKey string    `json:"licensekey"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Github struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username"`
	AccessToken  string             `json:"access_token"`
	RefreshToken string             `json:"refresh_token"`
	ExpiresIn    int                `json:"expires_in"`
	Type         string             `json:"token_type"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty" time_format:"2006-01-02 15:04:05" time_now:"true" default:"true" imutable:"true"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty" time_format:"2006-01-02 15:04:05" time_now:"true" default:"true"`
}
