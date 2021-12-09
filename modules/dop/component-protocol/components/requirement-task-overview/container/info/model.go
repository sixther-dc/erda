// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package info

import (
	"context"
	"encoding/json"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/modules/dop/dao"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
)

type Info struct {
	base.DefaultProvider

	Type   string          `json:"type"`
	Data   DataDetail      `json:"data"`
	Issues []dao.IssueItem `json:"-"`
}

type DataDetail struct {
	Data [][]Data `json:"data"`
}

type Data struct {
	Main string `json:"main"`
	Sub  string `json:"sub"`
}

func (i *Info) SetToProtocolComponent(c *cptype.Component) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &c)
}

func (i *Info) InitFromProtocol(ctx context.Context, c *cptype.Component) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, i)
}