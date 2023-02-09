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

package disk

import (
	dataproto "hcm/pkg/api/data-service/cloud/disk"
	"hcm/pkg/criteria/validator"
)

// DiskBaseCreateReq 云盘基础请求数据
type DiskBaseCreateReq struct {
	AccountID string  `json:"account_id" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Region    string  `json:"region" validate:"required"`
	Zone      string  `json:"zone" validate:"required"`
	DiskSize  uint64  `json:"disk_size" validate:"required"`
	DiskType  string  `json:"disk_type" validate:"required"`
	DiskCount uint32  `json:"disk_count" validate:"required"`
	Memo      *string `json:"memo"`
}

// DiskSyncReq disk sync request
type DiskSyncReq struct {
	AccountID         string `json:"account_id" validate:"required"`
	Region            string `json:"region" validate:"omitempty"`
	ResourceGroupName string `json:"resource_group_name" validate:"omitempty"`
}

// Validate disk sync request.
func (req *DiskSyncReq) Validate() error {
	return validator.Validate.Struct(req)
}

// DiskSyncDS disk data-service
type DiskSyncDS struct {
	IsUpdated bool
	HcDisk    *dataproto.DiskResult
}
