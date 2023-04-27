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

// Package subnet defines subnet service.
package subnet

import (
	subnetlogics "hcm/cmd/hc-service/logics/subnet"
	"hcm/pkg/adaptor/types"
	adcore "hcm/pkg/adaptor/types/core"
	"hcm/pkg/api/core"
	dataservice "hcm/pkg/api/data-service"
	"hcm/pkg/api/data-service/cloud"
	hcservice "hcm/pkg/api/hc-service"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/dal/dao/tools"
	"hcm/pkg/rest"
)

// AwsSubnetCreate create aws subnet.
func (s subnet) AwsSubnetCreate(cts *rest.Contexts) (interface{}, error) {
	req := new(hcservice.SubnetCreateReq[hcservice.AwsSubnetCreateExt])
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	awsCreateOpt := &subnetlogics.SubnetCreateOptions[hcservice.AwsSubnetCreateExt]{
		BkBizID:    req.BkBizID,
		AccountID:  req.AccountID,
		Region:     req.Extension.Region,
		CloudVpcID: req.CloudVpcID,
		CreateReqs: []hcservice.SubnetCreateReq[hcservice.AwsSubnetCreateExt]{*req},
	}
	res, err := s.subnet.AwsSubnetCreate(cts.Kit, awsCreateOpt)
	if err != nil {
		return nil, err
	}

	return core.CreateResult{ID: res.IDs[0]}, nil
}

// AwsSubnetUpdate update aws subnet.
func (s subnet) AwsSubnetUpdate(cts *rest.Contexts) (interface{}, error) {
	id := cts.PathParameter("id").String()

	req := new(hcservice.SubnetUpdateReq)
	if err := cts.DecodeInto(req); err != nil {
		return nil, errf.NewFromErr(errf.DecodeRequestFailed, err)
	}
	if err := req.Validate(); err != nil {
		return nil, errf.NewFromErr(errf.InvalidParameter, err)
	}

	getRes, err := s.cs.DataService().Aws.Subnet.Get(cts.Kit.Ctx, cts.Kit.Header(), id)
	if err != nil {
		return nil, err
	}

	cli, err := s.ad.Aws(cts.Kit, getRes.AccountID)
	if err != nil {
		return nil, err
	}

	updateOpt := new(types.AwsSubnetUpdateOption)
	err = cli.UpdateSubnet(cts.Kit, updateOpt)
	if err != nil {
		return nil, err
	}

	updateReq := &cloud.SubnetBatchUpdateReq[cloud.AwsSubnetUpdateExt]{
		Subnets: []cloud.SubnetUpdateReq[cloud.AwsSubnetUpdateExt]{{
			ID: id,
			SubnetUpdateBaseInfo: cloud.SubnetUpdateBaseInfo{
				Memo: req.Memo,
			},
		}},
	}
	err = s.cs.DataService().Aws.Subnet.BatchUpdate(cts.Kit.Ctx, cts.Kit.Header(), updateReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// AwsSubnetDelete delete aws subnet.
func (s subnet) AwsSubnetDelete(cts *rest.Contexts) (interface{}, error) {
	id := cts.PathParameter("id").String()

	getRes, err := s.cs.DataService().Aws.Subnet.Get(cts.Kit.Ctx, cts.Kit.Header(), id)
	if err != nil {
		return nil, err
	}

	cli, err := s.ad.Aws(cts.Kit, getRes.AccountID)
	if err != nil {
		return nil, err
	}

	delOpt := &adcore.BaseRegionalDeleteOption{
		BaseDeleteOption: adcore.BaseDeleteOption{ResourceID: getRes.CloudID},
		Region:           getRes.Region,
	}
	err = cli.DeleteSubnet(cts.Kit, delOpt)
	if err != nil {
		return nil, err
	}

	deleteReq := &dataservice.BatchDeleteReq{
		Filter: tools.EqualExpression("id", id),
	}
	err = s.cs.DataService().Global.Subnet.BatchDelete(cts.Kit.Ctx, cts.Kit.Header(), deleteReq)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// AwsSubnetCountIP count aws subnets' available ips.
func (s subnet) AwsSubnetCountIP(cts *rest.Contexts) (interface{}, error) {
	id := cts.PathParameter("id").String()
	if len(id) == 0 {
		return nil, errf.New(errf.InvalidParameter, "id is required")
	}

	getRes, err := s.cs.DataService().Aws.Subnet.Get(cts.Kit.Ctx, cts.Kit.Header(), id)
	if err != nil {
		return nil, err
	}

	cli, err := s.ad.Aws(cts.Kit, getRes.AccountID)
	if err != nil {
		return nil, err
	}

	listOpt := &adcore.AwsListOption{
		Region:   getRes.Region,
		CloudIDs: []string{getRes.CloudID},
	}
	subnetRes, err := cli.ListSubnet(cts.Kit, listOpt)
	if err != nil {
		return nil, err
	}

	if len(subnetRes.Details) != 1 {
		return nil, errf.New(errf.InvalidParameter, "subnet details count is invalid")
	}

	if subnetRes.Details[0].Extension == nil {
		return nil, errf.Newf(errf.InvalidParameter, "get aws subnet by cloud id %s failed", getRes.CloudID)
	}

	return &hcservice.SubnetCountIPResult{
		AvailableIPv4Count:  uint64(subnetRes.Details[0].Extension.AvailableIPAddressCount),
		TotalIpAddressCount: uint64(subnetRes.Details[0].Extension.TotalIpAddressCount),
		UsedIpAddressCount:  uint64(subnetRes.Details[0].Extension.UsedIpAddressCount),
	}, nil
}
