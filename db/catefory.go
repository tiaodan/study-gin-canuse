// db class 相关操作
package db

import (
	"study-spider-manhua-gin/logger"
	"study-spider-manhua-gin/models"

	// 导入 clause 包
	"gorm.io/gorm/clause"
)

// 增
func CategoryAdd(categoryData *models.Category) error {
	result := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "NameId"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": categoryData.Name}),
	}).Create(categoryData)
	if result.Error != nil {
		logger.Error("创建失败:", result.Error)
		return result.Error
	} else {
		logger.Debug("创建成功:", categoryData, "1")
	}
	return nil
}

// 批量增
func CategoriesBatchAdd(categories []*models.Category) {
	for i, categoryData := range categories {
		err := CategoryAdd(categoryData)
		if err == nil {
			logger.Debug("批量创建第%d条成功, category: %v", i+1, categoryData.Name)
		} else {
			logger.Error("批量创建第%d条失败, err: %v", i+1, err)
		}
	}
}

// 删
func CategoryDelete(id uint) {
	var categoryData models.Category
	result := DB.Delete(&categoryData, id)
	if result.Error != nil {
		logger.Error("删除失败:", result.Error)
	} else {
		logger.Debug("删除成功:", id)
	}
}

// 批量删
func CategoriesBatchDelete(ids []uint) {
	var categories []models.Category
	result := DB.Delete(&categories, ids)
	if result.Error != nil {
		logger.Error("批量删除失败:", result.Error)
	} else {
		logger.Debug("批量删除成功:", ids)
	}
}

// 改
func CategoryUpdate(nameId uint, updates map[string]interface{}) {
	var categoryData models.Category
	result := DB.Model(&categoryData).Where("name_id = ?", nameId).Updates(updates)
	if result.Error != nil {
		logger.Error("修改失败:", result.Error)
	} else {
		logger.Debug("修改成功:", nameId)
	}
}

// 批量改
func CategoriesBatchUpdate(updates map[uint]map[string]interface{}) {
	for nameId, update := range updates {
		var categoryData models.Category
		result := DB.Model(&categoryData).Where("name_id = ?", nameId).Updates(update)
		if result.Error != nil {
			logger.Error("更新类型 %d 失败: %v", nameId, result.Error)
		} else {
			logger.Debug("更新类型 %d 成功", nameId)
		}
	}
}

// 查
func CategoryQueryById(id uint) *models.Category {
	var categoryData models.Category
	result := DB.First(&categoryData, id)
	if result.Error != nil {
		logger.Error("查询失败:", result.Error)
		return nil
	}
	logger.Debug("查询成功:", categoryData)
	return &categoryData
}

// 批量查
func CategoriesBatchQuery(ids []uint) ([]*models.Category, error) {
	var categories []*models.Category
	result := DB.Find(&categories, ids)
	if result.Error != nil {
		logger.Error("批量查询失败: %v", result.Error)
		return categories, result.Error
	}
	logger.Debug("批量查询成功, 查询到 %d 条记录", len(categories))
	return categories, nil
}
