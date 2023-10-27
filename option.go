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

import "github.com/spf13/pflag"

type Options struct {
	Host   string `json:"host,omitempty"`
	Url    string `json:"url,omitempty"`
	Token  string `json:"token,omitempty"`
	Secret string `json:"secret,omitempty"`
	Level  int    `json:"level,omitempty"`
}

const (
	DefaultDingHost   = "https://oapi.dingtalk.com"
	DefaultDingUrl    = "/robot/send"
	DefaultDingToken  = "76f509430610108713fc835cf6e4cd460856c2e8d49f75e9e31bd04796be5f21"
	DefaultDingSecret = "SECed7b62be679450c4484799ea52d8cea2304e0b3b94b0618c0d587d45dd1f2c24"
	DefaultLevel      = 0
)

func NewOptions() *Options {
	return &Options{
		Host:   DefaultDingHost,
		Url:    DefaultDingUrl,
		Token:  DefaultDingToken,
		Secret: DefaultDingSecret,
		Level:  DefaultLevel,
	}
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "ding-host", o.Host, "dingding service host")
	fs.StringVar(&o.Url, "ding-url", o.Url, "dingding service robot token")
	fs.StringVar(&o.Token, "ding-token", o.Token, "dingding service robot token")
	fs.StringVar(&o.Secret, "ding-secret", o.Secret, "dingding service robot secret")
	fs.IntVar(&o.Level, "alert-level", o.Level, "dingding service robot level")
}
