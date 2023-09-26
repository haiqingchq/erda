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

package dop

import "github.com/erda-project/erda/internal/core/openapi/legacy/api/apis"

var FuncPointsQuery = apis.ApiSpec{
	Path:          "/api/func-points-trend/actions/query",
	BackendPath:   "/api/func-points-trend/actions/query",
	Host:          "dop.marathon.l4lb.thisdcos.directory:9527",
	Scheme:        "http",
	Method:        "POST",
	TryCheckLogin: true,
	Doc:           "summary: 功能点生产率PDR查询",
}