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

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sumengzs/alerter"
	"gopkg.in/resty.v1"
	"time"
)

func NewWithOptions(opts *Options) alerter.Alerter {
	c := &client{
		url:    opts.Url,
		token:  opts.Token,
		secret: opts.Secret,
		level:  opts.Level,
		client: resty.New().SetTimeout(time.Second * 5).SetHostURL(opts.Host),
	}
	return alerter.New(c)
}

func New() alerter.Alerter {
	return alerter.New(Client)
}

type client struct {
	url    string
	token  string
	secret string
	level  int
	client *resty.Client
}

var Client *client

func NewClient(opts *Options) {
	Client = &client{
		url:    opts.Url,
		token:  opts.Token,
		secret: opts.Secret,
		level:  opts.Level,
		client: resty.New().SetTimeout(time.Second * 5).SetHostURL(opts.Host),
	}
}

func (c *client) Enabled(level int) bool {
	return level >= c.level
}

func (c *client) Info(level int, msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c *client) Error(err error, msg string, keysAndValues ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (c *client) WithValues(keysAndValues ...interface{}) alerter.Sink {
	//TODO implement me
	panic("implement me")
}

func (c *client) WithName(name string) alerter.Sink {
	//TODO implement me
	panic("implement me")
}

func (c *client) Send(msg *Content) error {
	if len(c.token) == 0 {
		return fmt.Errorf("not access token")
	}
	req := c.client.R().
		SetBody(msg).
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetQueryParam("access_token", c.token)
	if len(c.secret) != 0 {
		t := time.Now().UnixNano() / 1e6
		req.SetQueryParam("timestamp", fmt.Sprintf("%d", t))
		req.SetQueryParam("sign", sign(t, c.secret))
	}
	resp, err := req.Post(c.url)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("send ding talk failed code[%d]", resp.StatusCode())
	}
	respMsg := &ResponseMsg{}
	err = json.Unmarshal(resp.Body(), respMsg)
	if err != nil {
		return err
	}
	if respMsg.ErrCode != 0 {
		return fmt.Errorf("send ding talk failed: %+v", *respMsg)
	}
	return nil
}

func sign(t int64, secret string) string {
	strToHash := fmt.Sprintf("%d\n%s", t, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}
