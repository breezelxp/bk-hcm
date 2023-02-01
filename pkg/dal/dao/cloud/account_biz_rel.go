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

package cloud

import (
	"fmt"

	"hcm/pkg/api/core"
	"hcm/pkg/criteria/errf"
	"hcm/pkg/dal/dao/orm"
	"hcm/pkg/dal/dao/tools"
	"hcm/pkg/dal/dao/types"
	"hcm/pkg/dal/table"
	"hcm/pkg/dal/table/cloud"
	"hcm/pkg/kit"
	"hcm/pkg/logs"
	"hcm/pkg/runtime/filter"

	"github.com/jmoiron/sqlx"
)

// AccountBizRel only used for account and biz rel.
type AccountBizRel interface {
	BatchCreateWithTx(kt *kit.Kit, tx *sqlx.Tx, rels []*cloud.AccountBizRelTable) error
	List(kt *kit.Kit, opt *types.ListOption) (*types.ListAccountBizRelDetails, error)
	DeleteWithTx(kt *kit.Kit, tx *sqlx.Tx, filterExpr *filter.Expression) error
}

var _ AccountBizRel = new(AccountBizRelDao)

// AccountBizRelDao account and biz relation dao.
type AccountBizRelDao struct {
	Orm orm.Interface
}

// BatchCreateWithTx AccountBizRel with tx.
func (a AccountBizRelDao) BatchCreateWithTx(kt *kit.Kit, tx *sqlx.Tx, rels []*cloud.AccountBizRelTable) error {
	if len(rels) == 0 {
		return errf.New(errf.InvalidParameter, "account_biz_rel is required")
	}

	sql := fmt.Sprintf(`INSERT INTO %s (%s)	VALUES(%s)`, table.AccountBizRelTable,
		cloud.AccountBizRelColumns.ColumnExpr(), cloud.AccountBizRelColumns.ColonNameExpr())

	err := a.Orm.Txn(tx).BulkInsert(kt.Ctx, sql, rels)
	if err != nil {
		return fmt.Errorf("insert %s failed, err: %v", table.AccountBizRelTable, err)
	}

	return nil
}

// List AccountBizRel list.
func (a AccountBizRelDao) List(kt *kit.Kit, opt *types.ListOption) (*types.ListAccountBizRelDetails, error) {
	if opt == nil {
		return nil, errf.New(errf.InvalidParameter, "list account_biz_rel options is nil")
	}

	if err := opt.Validate(filter.NewExprOption(filter.RuleFields(cloud.AccountBizRelColumns.ColumnTypes())),
		core.DefaultPageOption); err != nil {
		return nil, err
	}

	whereExpr, err := opt.Filter.SQLWhereExpr(tools.DefaultSqlWhereOption)
	if err != nil {
		return nil, err
	}

	if opt.Page.Count {
		// this is a count request, then do count operation only.
		sql := fmt.Sprintf(`SELECT COUNT(*) FROM %s %s`, table.AccountBizRelTable, whereExpr)

		count, err := a.Orm.Do().Count(kt.Ctx, sql)
		if err != nil {
			logs.ErrorJson("count account_biz_rel failed, err: %v, filter: %s, rid: %s", err, opt.Filter, kt.Rid)
			return nil, err
		}

		return &types.ListAccountBizRelDetails{Count: count}, nil
	}

	pageExpr, err := types.PageSQLExpr(opt.Page, types.DefaultPageSQLOption)
	if err != nil {
		return nil, err
	}

	sql := fmt.Sprintf(`SELECT %s FROM %s %s %s`, cloud.AccountBizRelColumns.FieldsNamedExpr(opt.Fields),
		table.AccountBizRelTable, whereExpr, pageExpr)

	details := make([]*cloud.AccountBizRelTable, 0)
	if err = a.Orm.Do().Select(kt.Ctx, &details, sql); err != nil {
		return nil, err
	}

	return &types.ListAccountBizRelDetails{Count: 0, Details: details}, nil
}

// DeleteWithTx AccountBizRel with tx.
func (a AccountBizRelDao) DeleteWithTx(kt *kit.Kit, tx *sqlx.Tx, filterExpr *filter.Expression) error {
	if filterExpr == nil {
		return errf.New(errf.InvalidParameter, "filter expr is required")
	}

	whereExpr, err := filterExpr.SQLWhereExpr(tools.DefaultSqlWhereOption)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf(`DELETE FROM %s %s`, table.AccountBizRelTable, whereExpr)
	if err := a.Orm.Txn(tx).Delete(kt.Ctx, sql); err != nil {
		logs.ErrorJson("delete account_biz_rel failed, err: %v, filter: %s, rid: %s", err, filterExpr, kt.Rid)
		return err
	}

	return nil
}
