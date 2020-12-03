package goGrowiAPI

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const GET_ENDPOINT string = "/_api/pages.get"

type PagesService service

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

func (p *PagesService) Get(ctx context.Context, path string) (*Page, error) {
	params := url.Values{}
	params.Set("access_token", p.client.config.Token)
	params.Set("path", path)

	res, err := p.client.newRequest(ctx, http.MethodGet, GET_ENDPOINT, &params)
	if err != nil {
		return nil, err
	}
	pagesGet := PagesGet{}
	if err := json.Unmarshal(res, &pagesGet); err != nil {
		return nil, err
	}

	if !pagesGet.Ok {
		return nil, fmt.Errorf("failed %+v", pagesGet)
	}

	return &pagesGet.Page, nil
}
