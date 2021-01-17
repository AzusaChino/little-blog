package easy_orm

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// sumRes 求和
type sumRes struct {
	Total int64 `json:"total"`
}

// Field 查询字段结构体
type Field struct {
	Condition string      `json:"condition"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"` // int  []int uint  []uint string []string
}

// Relate 关联关系
type Relate struct {
	Value string
	Func  interface{}
}

// Search 查询参数结构体
type Search struct {
	Selects   []string  `json:"selects"`
	Fields    []*Field  `json:"fields"`
	Relations []*Relate `json:"relations"`
	OrderBy   string    `json:"order_by"`
	Sort      string    `json:"sort"`
	Limit     int       `json:"limit"`
	Offset    int       `json:"offset"`
}

// Count field
func Count(model interface{}, field string) (int64, error) {
	var sr sumRes
	s := fmt.Sprintf("sum(`%s`) as total", field)
	err := EzOrm.Db.Model(model).Select(s).Scan(&sr).Error
	if err != nil {
		return 0, err
	}
	return sr.Total, nil
}

// getAll 批量查询
func getAll(model interface{}, s *Search) *gorm.DB {
	sort := "desc"
	orderBy := "created_at"
	if len(s.Sort) > 0 {
		sort = s.Sort
	}
	if len(s.OrderBy) > 0 {
		orderBy = s.OrderBy
	}

	db := EzOrm.Db.Model(model).
		Order(fmt.Sprintf("%s %s", orderBy, sort)).
		Scopes(FoundByWhereScope(s.Fields), RelationScope(s.Relations))

	return db
}

// Paginate 分页查询
func Paginate(model, data interface{}, s *Search) (int64, error) {
	var count int64
	db := getAll(model, s)
	if err := db.Count(&count).Error; err != nil {
		return count, err
	}
	db = db.Scopes(PaginateScope(s.Offset, s.Limit))

	if err := db.Select(s.Selects).Find(data).Error; err != nil {
		return count, err
	}

	return count, nil
}

// All 批量查询
func All(model, data interface{}, s *Search) error {
	if err := getAll(model, s).Select(s.Selects).Find(data).Error; err != nil {
		return err
	}
	return nil
}

// First
func First(model interface{}, search *Search) error {
	err := Found(search).First(model).Error
	if err != nil {
		return err
	}
	return nil
}

// FindById
func FindById(model interface{}, id uint) error {
	err := EzOrm.Db.First(model, id).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除
func Delete(model interface{}, s *Search) error {
	if err := Found(s).Delete(model).Error; err != nil {
		return err
	}
	return nil
}

// Delete 通过 id 删除
func DeleteById(model interface{}, id uint) error {
	if err := EzOrm.Db.Delete(model, id).Error; err != nil {
		return err
	}
	return nil
}

// Create 新建
func Create(model interface{}) error {
	if err := EzOrm.Db.Create(model).Error; err != nil {
		return err
	}
	return nil
}

// Save 保存
func Save(model interface{}) error {
	if err := EzOrm.Db.Save(model).Error; err != nil {
		return err
	}
	return nil
}

// Found 查询条件
func Found(s *Search) *gorm.DB {
	return EzOrm.Db.Scopes(RelationScope(s.Relations), FoundByWhereScope(s.Fields)).Select(s.Selects)
}

// Update 更新
func Update(v, d interface{}, fileds []interface{}, id uint) error {
	u := EzOrm.Db.Model(v).Where("id = ?", id)
	if len(fileds) > 0 {
		if err := u.Select(fileds[0], fileds[1:]...).Updates(d).Error; err != nil {
			return err
		}
	} else {
		if err := u.Updates(d).Error; err != nil {
			return err
		}
	}

	return nil
}

// UpdateWithFilde 更新
func UpdateWithFilde(v interface{}, filed map[string]interface{}, id uint) error {
	if err := EzOrm.Db.Model(v).Where("id = ?", id).Updates(filed).Error; err != nil {
		return err
	}

	return nil
}

// RelationScope 加载关联关系
func RelationScope(relates []*Relate) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(relates) > 0 {
			for _, re := range relates {
				if len(re.Value) > 0 {
					if re.Func != nil {
						db = db.Preload(re.Value, re.Func)
					} else {
						db = db.Preload(re.Value)
					}
				}
			}
		}
		return db
	}
}

// FoundByWhereScope 查询条件
func FoundByWhereScope(fields []*Field) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(fields) > 0 {
			for _, field := range fields {
				if field != nil {
					if field.Condition == "" {
						field.Condition = "="
					}
					if value, ok := field.Value.(int); ok {
						if value > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(uint); ok {
						if value > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.(string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]int); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]uint); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else if value, ok := field.Value.([]string); ok {
						if len(value) > 0 {
							db = db.Where(fmt.Sprintf("%s %s ?", field.Key, field.Condition), value)
						}
					} else {
						//i := field.Value
						color.Red(fmt.Sprintf("未知数据类型：%+v ", field.Value))
					}
				}
			}
		}
		return db
	}
}

// GetRelations 转换前端获取关联关系为 []*Relate
func GetRelations(relation string, fs map[string]interface{}) []*Relate {
	var relates []*Relate
	if len(relation) > 0 {
		res := strings.Split(relation, ",")
		for _, re := range res {
			relate := &Relate{
				Value: re,
			}
			// 增加关联过滤
			for key, f := range fs {
				if key == re {
					relate.Func = f
				}
			}
			relates = append(relates, relate)
		}
	}
	return relates
}

// GetFields 转换前端查询关系为 []*Field
func GetFields(searchs map[string]interface{}) []*Field {
	var fields []*Field
	for key, search := range searchs {
		field := GetField(key, search)
		fields = append(fields, field)
	}
	return fields
}

// GetSelects 字段过滤 []string
func GetSelects(field string) []string {
	if field == "" {
		return nil
	}
	return strings.Split(field, ",")
}

// GetField 转换前端查询关系为 *Field
func GetField(key string, search interface{}) *Field {
	if s, ok := search.(string); ok {
		if len(s) > 0 {
			if strings.Contains(s, ":") {
				searches := strings.Split(s, ":")
				if len(searches) == 2 {
					value := searches[0]
					if strings.ToLower(searches[1]) == "like" {
						value = fmt.Sprintf("%%%s%%", searches[0])
					}

					return &Field{
						Condition: searches[1],
						Key:       key,
						Value:     value,
					}

				} else if len(searches) == 1 {
					return &Field{
						Condition: "=",
						Key:       key,
						Value:     searches[0],
					}
				}
			} else {
				return &Field{
					Condition: "=",
					Key:       key,
					Value:     search,
				}
			}
		}
	} else { // int string []int []string
		return &Field{
			Condition: "=",
			Key:       key,
			Value:     search,
		}
	}

	return nil
}

// PaginateScope 分页
func PaginateScope(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize < 0:
			pageSize = -1
		case pageSize == 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		if page < 0 {
			offset = -1
		}
		return db.Offset(offset).Limit(pageSize)
	}
}

// IsNotFound 判断是否是查询不存在错误
func IsNotFound(err error) bool {
	if ok := errors.Is(err, gorm.ErrRecordNotFound); ok {
		color.Yellow("查询数据不存在")
		return true
	}
	return false
}

// GetRolesForUser 获取角色
func GetRolesForUser(uid uint) []string {
	uids, err := EzOrm.Enforcer.GetRolesForUser(strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		color.Red(fmt.Sprintf("GetRolesForUser 错误: %v", err))
		return []string{}
	}

	return uids
}

// GetPermissionsForUser 获取角色权限
func GetPermissionsForUser(uid uint) [][]string {
	return EzOrm.Enforcer.GetPermissionsForUser(strconv.FormatUint(uint64(uid), 10))
}
