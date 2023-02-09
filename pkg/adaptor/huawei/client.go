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

package huawei

import (
	"fmt"

	"hcm/pkg/adaptor/types"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	evs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/evs/v2"
	iam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iamregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
	vpcv2 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3"
	vpcregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v3/region"
)

type clientSet struct {
	credentials *basic.Credentials
}

func newClientSet(secret *types.BaseSecret) *clientSet {
	return &clientSet{
		credentials: basic.NewCredentialsBuilder().
			WithAk(secret.CloudSecretID).
			WithSk(secret.CloudSecretKey).
			Build(),
	}
}

func (c *clientSet) iamClient(region *region.Region) (*iam.IamClient, error) {
	client := iam.NewIamClient(
		iam.IamClientBuilder().
			WithRegion(region).
			WithCredential(c.credentials).
			WithHttpConfig(config.DefaultHttpConfig()).
			Build())

	return client, nil
}

func (c *clientSet) iamRegionClient(region string) (*iam.IamClient, error) {
	client := iam.NewIamClient(
		iam.IamClientBuilder().
			WithRegion(iamregion.ValueOf(region)).
			WithCredential(c.credentials).
			WithHttpConfig(config.DefaultHttpConfig()).
			Build())

	return client, nil
}

func (c *clientSet) evsClient(region *region.Region) (*evs.EvsClient, error) {
	client := evs.NewEvsClient(
		evs.EvsClientBuilder().
			WithRegion(region).
			WithCredential(c.credentials).
			WithHttpConfig(config.DefaultHttpConfig()).
			Build())

	return client, nil
}

func (c *clientSet) vpcClient(regionID string) (cli *vpc.VpcClient, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic recovered, err: %v", p)
		}
	}()

	client := vpc.NewVpcClient(
		vpc.VpcClientBuilder().
			WithRegion(vpcregion.ValueOf(regionID)).
			WithCredential(c.credentials).
			WithHttpConfig(config.DefaultHttpConfig()).
			Build())

	return client, nil
}

func (c *clientSet) vpcClientV2(regionID string) (cli *vpcv2.VpcClient, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic recovered, err: %v", p)
		}
	}()

	client := vpcv2.NewVpcClient(
		vpcv2.VpcClientBuilder().
			WithRegion(vpcregion.ValueOf(regionID)).
			WithCredential(c.credentials).
			WithHttpConfig(config.DefaultHttpConfig()).
			Build())

	return client, nil
}
