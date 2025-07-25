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

package admin

type ProviderVO struct {
	ID     int64     `json:"id"`
	Name   string    `json:"name"`
	ApiKey string    `json:"apiKey,omitempty"`
	Models []ModelVO `json:"models"`
}

type ModelVO struct {
	ID          int64  `json:"id"`
	Pid         int64  `json:"pid"`
	Name        string `json:"name"`
	InputPrice  int64  `json:"inputPrice"`
	OutputPrice int64  `json:"outputPrice"`
	PriceMode   string `json:"price_mode"`
}
