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

package producer

import (
	"errors"

	"hcm/pkg/async/action"
	"hcm/pkg/async/backend"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/criteria/validator"
	tableasync "hcm/pkg/dal/table/async"
	"hcm/pkg/dal/table/types"
)

// AddTemplateFlowOption define add flow option.
type AddTemplateFlowOption struct {
	// Name 任务流模版名称
	Name enumor.FlowName `json:"name" validate:"required"`
	// Memo 备注
	Memo string `json:"memo" validate:"omitempty"`
	// Tasks 任务私有化参数设置
	Tasks []TemplateFlowTask `json:"tasks" validate:"omitempty"`
	// IsInitState 是否初始化状态
	IsInitState bool `json:"is_init_state" validate:"omitempty"`
}

// Validate AddTemplateFlowOption
func (opt *AddTemplateFlowOption) Validate() error {

	if err := validator.Validate.Struct(opt); err != nil {
		return err
	}

	if err := opt.Name.Validate(); err != nil {
		return err
	}

	for index := range opt.Tasks {
		if err := opt.Tasks[index].Validate(); err != nil {
			return err
		}
	}

	return nil
}

// TemplateFlowTask define task info.
type TemplateFlowTask struct {
	// ActionID 任务在当前任务流模版中的唯一ID
	ActionID action.ActIDType `json:"action_id" validate:"required"`
	// Params 任务执行请求参数
	Params types.JsonField `json:"params" validate:"required"`
}

// Validate TemplateFlowTask
func (task *TemplateFlowTask) Validate() error {
	return validator.Validate.Struct(task)
}

// AddCustomFlowOption define add custom flow option.
type AddCustomFlowOption struct {
	// Name 任务流模版名称
	Name enumor.FlowName `json:"name" validate:"required"`
	// Memo 备注
	Memo string `json:"memo" validate:"omitempty"`
	// ShareData 共享数据
	ShareData *tableasync.ShareData `json:"share_data" validate:"omitempty"`
	// Tasks 任务私有化参数设置
	Tasks []CustomFlowTask `json:"tasks" validate:"required"`
	// IsInitState 是否初始化状态
	IsInitState bool `json:"is_init_state" validate:"omitempty"`
}

// Validate AddCustomFlowOption
func (opt *AddCustomFlowOption) Validate() error {

	if err := opt.Name.Validate(); err != nil {
		return err
	}

	if len(opt.Tasks) == 0 {
		return errors.New("tasks is required")
	}

	for _, task := range opt.Tasks {
		if err := task.Validate(); err != nil {
			return err
		}
	}

	return validator.Validate.Struct(opt)
}

// CustomFlowTask define custom flow task info.
type CustomFlowTask struct {
	// ActionID Action唯一序列号
	ActionID action.ActIDType `json:"action_id" validate:"required"`
	// ActionName Action名称
	ActionName enumor.ActionName `json:"action_name" validate:"required"`
	// DependOn 运行当前Action依赖的前置ActionID
	DependOn []action.ActIDType `json:"depend_on" validate:"omitempty"`
	// Params 执行请求参数
	Params types.JsonField `json:"params" validate:"omitempty"`
	// Retry 任务运行重试相关配置参数，如果不设置，默认不允许进行重试。
	Retry *tableasync.Retry `json:"retry" validate:"omitempty"`
}

// Validate CustomFlowTask
func (task *CustomFlowTask) Validate() error {
	if err := validator.Validate.Struct(task); err != nil {
		return err
	}

	if err := task.ActionName.Validate(); err != nil {
		return err
	}

	return nil
}

// UpdateCustomFlowStateOption define update custom flow state option.
type UpdateCustomFlowStateOption struct {
	// FlowInfos 任务状态信息
	FlowInfos []backend.UpdateFlowInfo `json:"flow_infos" validate:"required,min=1"`
}

// Validate UpdateCustomFlowStateOption
func (opt *UpdateCustomFlowStateOption) Validate() error {
	if len(opt.FlowInfos) == 0 {
		return errors.New("flow_infos is required")
	}

	for _, flowInfo := range opt.FlowInfos {
		if err := flowInfo.Validate(); err != nil {
			return err
		}
	}

	return validator.Validate.Struct(opt)
}
