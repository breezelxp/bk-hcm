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

package eip

import (
	"hcm/pkg/criteria/validator"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/eip/v2/model"
)

// HuaWeiEipListOption ...
type HuaWeiEipListOption struct {
	Region string `validate:"required"`
	Limit  *int32
	Marker *string
}

// Validate ...
func (o *HuaWeiEipListOption) Validate() error {
	return validator.Validate.Struct(o)
}

// ToListPublicipsRequest ...
func (o *HuaWeiEipListOption) ToListPublicipsRequest() (*model.ListPublicipsRequest, error) {
	req := &model.ListPublicipsRequest{Limit: o.Limit, Marker: o.Marker}
	return req, nil
}

// HuaWeiEipListResult ...
type HuaWeiEipListResult struct {
	Details []*HuaWeiEip
}

// HuaWeiEip ...
type HuaWeiEip struct {
	CloudID       string
	Name          *string
	Region        string
	InstanceId    *string
	Status        *string
	PublicIp      *string
	PrivateIp     *string
	PortID        *string
	BandwidthId   *string
	BandwidthName *string
	BandwidthSize *int32
}
