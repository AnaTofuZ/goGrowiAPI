package goGrowiAPI

import "time"

type User struct {
	IsGravatarEnabled bool      `json:"isGravatarEnabled"`
	IsEmailPublished  bool      `json:"isEmailPublished"`
	Lang              string    `json:"lang"`
	Status            int       `json:"status"`
	Admin             bool      `json:"admin"`
	ID                string    `json:"_id"`
	CreatedAt         time.Time `json:"createdAt"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	LastLoginAt       time.Time `json:"lastLoginAt"`
	ImageURLCached    string    `json:"imageUrlCached"`
}

type Page struct {
	Status         string        `json:"status"`
	Grant          int           `json:"grant"`
	GrantedUsers   []interface{} `json:"grantedUsers"`
	Liker          []interface{} `json:"liker"`
	SeenUsers      []string      `json:"seenUsers"`
	CommentCount   int           `json:"commentCount"`
	Extended       string        `json:"extended"`
	SubID          string        `json:"_id"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	Path           string        `json:"path"`
	Creator        User          `json:"creator"`
	LastUpdateUser User          `json:"lastUpdateUser"`
	RedirectTo     interface{}   `json:"redirectTo"`
	GrantedGroup   interface{}   `json:"grantedGroup"`
	V              int           `json:"__v"`
	Revision       struct {
		Format        string    `json:"format"`
		ID            string    `json:"_id"`
		CreatedAt     time.Time `json:"createdAt"`
		Path          string    `json:"path"`
		Body          string    `json:"body"`
		Author        User      `json:"author"`
		HasDiffToPrev bool      `json:"hasDiffToPrev"`
		V             int       `json:"__v"`
	} `json:"revision"`
	ID string `json:"id"`
}

type PagesGet struct {
	Page  `json:"page"`
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}
