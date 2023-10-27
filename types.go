/*
Copyright 2023 The dingtalkalerter Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dingtalkalerter

type MsgType string

const (
	TEXT        MsgType = "text"
	LINK        MsgType = "link"
	MARKDOWN    MsgType = "markdown"
	ACTION_CARD MsgType = "actionCard"
	FEED_CARD   MsgType = "feedCard"
)

type Content struct {
	Type       MsgType     `json:"msgtype"`
	Text       *Text       `json:"text,omitempty"`
	Markdown   *Markdown   `json:"markdown,omitempty"`
	Link       *Link       `json:"link,omitempty"`
	ActionCard *ActionCard `json:"actionCard,omitempty"`
	FeedCard   *FeedCard   `json:"feedCard,omitempty"`
	At         *At         `json:"at,omitempty"`
}

type Text struct {
	Content string `json:"content,omitempty"`
}

type Markdown struct {
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

type Link struct {
	Text       string `json:"text,omitempty"`
	Title      string `json:"title,omitempty"`
	PicUrl     string `json:"picUrl,omitempty"`
	MessageUrl string `json:"messageUrl,omitempty"`
}

type ActionCard struct {
	Title          string                `json:"title,omitempty"`
	Text           string                `json:"text,omitempty"`
	BtnOrientation ButtonOrientationType `json:"btnOrientation,omitempty"`
	SingleTitle    string                `json:"singleTitle,omitempty"`
	SingleURL      string                `json:"singleURL,omitempty"`
	Buttons        []Button              `json:"btns,omitempty"`
}

type ButtonOrientationType string

const (
	HORIZONTAL ButtonOrientationType = "0" // 横向
	VERTICAL   ButtonOrientationType = "1" // 竖向
)

type Button struct {
	Title     string `json:"title,omitempty"`
	ActionURL string `json:"actionURL,omitempty"`
}

type FeedCard struct {
	Links []FeedCardLink `json:"links,omitempty"`
}

type FeedCardLink struct {
	Title      string `json:"title,omitempty"`
	MessageURL string `json:"messageURL,omitempty"`
	PicURL     string `json:"picURL,omitempty"`
}

type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	AtUserIds []string `json:"atUserIds,omitempty"`
	IsAtAll   bool     `json:"isAtAll,omitempty"`
}

type ResponseMsg struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
