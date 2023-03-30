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
	"net/http"

	"hcm/cmd/hc-service/service/capability"
	"hcm/pkg/rest"
)

// InitEipService initial the eip service
func InitEipService(cap *capability.Capability) {
	e := &eipAdaptor{
		adaptor: cap.CloudAdaptor,
		dataCli: cap.ClientSet.DataService(),
	}

	h := rest.NewHandler()

	// 删除 Eip
	h.Add("DeleteEip", http.MethodDelete, "/vendors/{vendor}/eips", e.DeleteEip)
	// 关联 Eip
	h.Add("AssociateEip", http.MethodPost, "/vendors/{vendor}/eips/associate", e.AssociateEip)
	// 解关联 Eip
	h.Add("DisassociateEip", http.MethodPost, "/vendors/{vendor}/eips/disassociate", e.DisassociateEip)
	// 创建 Eip
	h.Add("CreateEip", http.MethodPost, "/vendors/{vendor}/eips/create", e.CreateEip)

	h.Load(cap.WebService)
}
