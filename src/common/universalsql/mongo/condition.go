/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mongo

import (
	"configcenter/src/common/mapstr"
	"configcenter/src/common/universalsql"
)

var _ universalsql.Condition = (*condition)(nil)

type condition struct {
}

func (cond *condition) ToSQL() string {
	return ""
}

func (cond *condition) ToMapStr() mapstr.MapStr {
	return nil
}

func (cond *condition) Element(element universalsql.ConditionElement) universalsql.Condition {
	return nil
}

func (cond *condition) And(elements ...universalsql.ConditionElement) universalsql.Condition {
	return nil
}

func (cond *condition) Or(elements ...universalsql.ConditionElement) universalsql.Condition {
	return nil
}

func (cond *condition) Embed(embedName string) universalsql.Condition {
	return nil
}