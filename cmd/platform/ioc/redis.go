// Copyright 2021 ecodeclub
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
	"fmt"
	"github.com/gotomicro/ego/core/econf"
	"github.com/redis/go-redis/v9"
)

func InitRedis() redis.Cmdable {
	type RedisConfig struct {
		Addr string `yaml:"addr"`
	}

	var cfg RedisConfig
	err := econf.UnmarshalKey("redis", &cfg)
	if err != nil {
		panic(fmt.Errorf("初始化 Redis 失败 %w", err))
	}

	return redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
}
