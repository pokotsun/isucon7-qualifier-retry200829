package main

import "time"

type User struct {
	ID          int64     `json:"-" db:"id"`
	Name        string    `json:"name" db:"name"`
	Salt        string    `json:"-" db:"salt"`
	Password    string    `json:"-" db:"password"`
	DisplayName string    `json:"display_name" db:"display_name"`
	AvatarIcon  string    `json:"avatar_icon" db:"avatar_icon"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
}

type Message struct {
	ID        int64     `db:"id"`
	ChannelID int64     `db:"channel_id"`
	UserID    int64     `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type ChannelInfo struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	UpdatedAt   time.Time `db:"updated_at"`
	CreatedAt   time.Time `db:"created_at"`
}

type HaveRead struct {
	UserID    int64     `db:"user_id"`
	ChannelID int64     `db:"channel_id"`
	MessageID int64     `db:"message_id"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

// getMessageのため
type MessageUser struct {
	ID        int64     `db:"id"`
	ChannelID int64     `db:"channel_id"`
	UserID    int64     `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`

	Name        string `json:"name" db:"name"`
	DisplayName string `json:"display_name" db:"display_name"`
	AvatarIcon  string `json:"avatar_icon" db:"avatar_icon"`
}

// IconのInitializeのため
type Icon struct {
	Name string `db:"name"`
	Data []byte `db:"data"`
}

type ChannelHaveRead struct {
	ID        int64 `db:"id"`
	MessageID int64 `db:"message_id"`
}

type ChannelCount struct {
	ChannelID int64 `db:"channel_id"`
	Count     int64 `db:"cnt"`
}
