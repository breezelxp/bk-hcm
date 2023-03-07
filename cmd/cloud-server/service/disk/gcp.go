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
	cloudproto "hcm/pkg/api/cloud-server/disk"
	hcproto "hcm/pkg/api/hc-service/disk"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/rest"
)

func (svc *diskSvc) gcpAttachDisk(cts *rest.Contexts) (interface{}, error) {
	req := new(cloudproto.GcpDiskAttachReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	// TODO 增加鉴权和审计
	// TODO 判断云盘是否可挂载

	basicInfo, err := svc.client.DataService().Global.Cloud.GetResourceBasicInfo(
		cts.Kit.Ctx,
		cts.Kit.Header(),
		enumor.DiskCloudResType,
		req.DiskID,
	)
	if err != nil {
		return nil, err
	}

	return nil, svc.client.HCService().Gcp.Disk.AttachDisk(
		cts.Kit.Ctx,
		cts.Kit.Header(),
		&hcproto.GcpDiskAttachReq{
			AccountID: basicInfo.AccountID,
			CvmID:     req.CvmID,
			DiskID:    req.DiskID,
		},
	)
}
