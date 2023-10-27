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

type AtOption interface {
	apply(model *At)
}

type funcAtOption struct {
	f func(model *At)
}

func (fdo *funcAtOption) apply(do *At) {
	fdo.f(do)
}

func newFuncAtOption(f func(model *At)) *funcAtOption {
	return &funcAtOption{f: f}
}

func WithAtAll() AtOption {
	return newFuncAtOption(func(o *At) {
		o.IsAtAll = true
	})
}

func WithAtMobiles(mobiles []string) AtOption {
	return newFuncAtOption(func(o *At) {
		o.AtMobiles = mobiles
	})
}

func WithAtUsers(users []string) AtOption {
	return newFuncAtOption(func(o *At) {
		o.AtUserIds = users
	})
}

func NewText(data string, opts ...AtOption) *Content {
	c := &Content{
		Type: TEXT,
		Text: &Text{Content: data},
		At:   &At{},
	}
	for _, opt := range opts {
		opt.apply(c.At)
	}
	return c
}

func NewLink(title, text, picUrl, msgUrl string) *Content {
	return &Content{
		Type: LINK,
		Link: &Link{
			Text:       text,
			Title:      title,
			PicUrl:     picUrl,
			MessageUrl: msgUrl,
		},
	}
}

func NewMarkDown(title, text string, opts ...AtOption) *Content {
	c := &Content{
		Type: MARKDOWN,
		Markdown: &Markdown{
			Text:  text,
			Title: title,
		},
		At: &At{},
	}
	for _, opt := range opts {
		opt.apply(c.At)
	}
	return c
}
