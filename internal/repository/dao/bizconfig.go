// Copyright 2021 ecodeclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao

import (
	"time"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type BizConfig struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	OwnerID   int64  `gorm:"column:owner_id;type:bigint;not null"`
	OwnerType string `gorm:"column:owner_type;type:varchar(20);not null"`
	Config    string `gorm:"column:config;type:text"`
	Ctime     int64
	Utime     int64
}

func (BizConfig) TableName() string {
	return "biz_configs"
}

type BizConfigDAO struct {
	db *gorm.DB
}

func NewBizConfigDAO(db *gorm.DB) *BizConfigDAO {
	return &BizConfigDAO{db: db}
}

func (d *BizConfigDAO) Insert(ctx context.Context, bc *BizConfig) (BizConfig, error) {
	now := time.Now().UnixMilli()
	bc.Ctime = now
	bc.Utime = now
	err := d.db.WithContext(ctx).Create(bc).Error
	if err != nil {
		return BizConfig{}, err
	}
	return *bc, nil
}

// GetByID 根据ID查询
func (d *BizConfigDAO) GetByID(ctx context.Context, id int64) (BizConfig, error) {
	var bc BizConfig
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&bc).Error
	return bc, err
}

// Update 更新记录
func (d *BizConfigDAO) Update(ctx context.Context, bc *BizConfig) error {
	bc.Utime = time.Now().UnixMilli()
	return d.db.WithContext(ctx).Save(bc).Error
}

// Delete 删除记录
func (d *BizConfigDAO) Delete(ctx context.Context, id string) error {
	return d.db.WithContext(ctx).Where("id = ?", id).Delete(&BizConfig{}).Error
}
