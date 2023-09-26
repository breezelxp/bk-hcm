/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package instancetype

import (
	"hcm/pkg/criteria/validator"
	"hcm/pkg/rest"
)

// HuaWeiInstanceTypeListReq ...
type HuaWeiInstanceTypeListReq struct {
	AccountID string `json:"account_id" validate:"required"`
	Region    string `json:"region" validate:"required"`
	Zone      string `json:"zone" validate:"required"`
}

// Validate ...
func (req *HuaWeiInstanceTypeListReq) Validate() error {
	return validator.Validate.Struct(req)
}

// HuaWeiInstanceTypeResp ...
type HuaWeiInstanceTypeResp struct {
	InstanceType            string `json:"instance_type"`
	GPU                     int64  `json:"gpu"`
	CPU                     int64  `json:"cpu"`
	Memory                  int64  `json:"memory"`
	FPGA                    int64  `json:"fpga"`
	InstanceFamily          string `json:"instance_family"`
	Name                    string `json:"name"`
	EcsinstanceArchitecture string `json:"ecsinstance_architecture"`
	Infocpuname             string `josn:"infocpuname"`
	QuotamaxRate            string `json:"quotamax_rate"`
	QuotamaxPps             string `json:"quotamax_pps"`
}

// HuaWeiInstanceTypeListResp ...
type HuaWeiInstanceTypeListResp struct {
	rest.BaseResp `json:",inline"`
	Data          []*HuaWeiInstanceTypeResp `json:"data"`
}
