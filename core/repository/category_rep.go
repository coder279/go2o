/**
 * Copyright 2015 @ z3q.net.
 * name : category_repo.go
 * author : jarryliu
 * date : 2016-06-04 13:01
 * description :
 * history :
 */
package repository

import (
	"fmt"
	"github.com/jsix/gof/db"
	"github.com/jsix/gof/db/orm"
	"github.com/jsix/gof/storage"
	"go2o/core/domain/interface/sale"
	"go2o/core/domain/interface/valueobject"
	saleImpl "go2o/core/domain/sale"
	"sort"
)

var _ sale.ICategoryRepo = new(categoryRepo)

type categoryRepo struct {
	db.Connector
	_valRepo          valueobject.IValueRepo
	_globCateManager sale.ICategoryManager
	storage          storage.Interface
}

func NewCategoryRepo(conn db.Connector, valRepo valueobject.IValueRepo,
	storage storage.Interface) sale.ICategoryRepo {
	return &categoryRepo{
		Connector: conn,
		_valRepo:   valRepo,
		storage:   storage,
	}
}

func (c *categoryRepo) GetGlobManager() sale.ICategoryManager {
	if c._globCateManager == nil {
		c._globCateManager = saleImpl.NewCategoryManager(0, c, c._valRepo)
	}
	return c._globCateManager
}

func (c *categoryRepo) getCategoryCacheKey(id int32) string {
	return fmt.Sprintf("go2o:rep:cat:c%d", id)
}

func (c *categoryRepo) SaveCategory(v *sale.Category) (int32, error) {
	id, err := orm.I32(orm.Save(c.GetOrm(), v, int(v.Id)))
	// 清理缓存
	if err == nil {
		c.storage.Del(c.getCategoryCacheKey(id))
		PrefixDel(c.storage, "go2o:rep:cat:*")
	}
	return id, err
}

// 检查分类是否关联商品
func (c *categoryRepo) CheckGoodsContain(mchId, id int32) bool {
	num := 0
	//清理项
	c.Connector.ExecScalar(`SELECT COUNT(0) FROM gs_item WHERE category_id IN
		(SELECT Id FROM cat_category WHERE mch_id=? AND id=?)`, &num, mchId, id)
	return num > 0
}

func (c *categoryRepo) DeleteCategory(mchId, id int32) error {
	//删除子类
	_, _, err := c.Connector.Exec("DELETE FROM cat_category WHERE mch_id=? AND parent_id=?",
		mchId, id)

	//删除分类
	_, _, err = c.Connector.Exec("DELETE FROM cat_category WHERE mch_id=? AND id=?",
		mchId, id)

	// 清理缓存
	if err == nil {
		c.storage.Del(c.getCategoryCacheKey(id))
		PrefixDel(c.storage, "go2o:rep:cat:*")
	}

	return err
}

func (c *categoryRepo) GetCategory(mchId, id int32) *sale.Category {
	e := sale.Category{}
	key := c.getCategoryCacheKey(id)
	if c.storage.Get(key, &e) != nil {
		err := c.Connector.GetOrm().Get(id, &e)
		if err != nil {
			return nil
		}
		c.storage.Set(key, &e)
	}
	return &e
}

// 创建分类
func (c *categoryRepo) CreateCategory(v *sale.Category) sale.ICategory {
	return saleImpl.NewCategory(c, v)
}

func (c *categoryRepo) convertICategory(list sale.CategoryList) []sale.ICategory {
	sort.Sort(list)
	slice := make([]sale.ICategory, len(list))
	for i, v := range list {
		slice[i] = c.CreateCategory(v)
	}
	return slice
}

func (c *categoryRepo) redirectGetCats() []*sale.Category {
	list := []*sale.Category{}
	err := c.Connector.GetOrm().Select(&list, "")
	if err != nil {
		handleError(err)
	}
	return list
}

func (c *categoryRepo) GetCategories(mchId int32) []*sale.Category {
	return c.redirectGetCats()
	//todo: cache
	//key := fmt.Sprintf("go2o:rep:cat:list9:%d", mchId)
	//list := []*sale.Category{}
	//if err := c.storage.Get(key, &list);err != nil {
	//    handleError(err)
	//    err := c.Connector.GetOrm().Select(&list, "mch_id=? ORDER BY id ASC", mchId)
	//    if err == nil {
	//        c.storage.SetExpire(key,list, DefaultCacheSeconds)
	//    } else {
	//        handleError(err)
	//    }
	//}
	//return list
}
