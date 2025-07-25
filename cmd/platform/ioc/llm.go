// Copyright 2023 ecodeclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ioc

import (
	"github.com/ecodeclub/ai-gateway-go/internal/service/llm"
	"github.com/ecodeclub/ai-gateway-go/internal/service/llm/openai"
	"github.com/gotomicro/ego/core/econf"
)

func initLLMHandler() llm.Handler {
	type AliyunConfig struct {
		APIKey  string `json:"apiKey"`
		BaseURL string `json:"baseURL"`
		Model   string `json:"model"`
	}
	var cfg AliyunConfig

	err := econf.UnmarshalKey("aliyun", &cfg)
	if err != nil {
		panic(err)
	}
	handler := openai.NewHandler(cfg.APIKey, cfg.BaseURL, cfg.Model)
	return handler
}
