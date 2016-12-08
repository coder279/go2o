package promodel

import "go2o/core/domain/interface/pro_model"

var _ promodel.IBrandService

type BrandServiceImpl struct {
    rep promodel.IProBrandRepo
}

func NewBrandService(rep promodel.IProBrandRepo) *BrandServiceImpl {
    return &BrandServiceImpl{
        rep:rep,
    }
}

// 获取品牌
func (b *BrandServiceImpl)Get(brandId int32) *promodel.ProBrand {
    return b.rep.GetProBrand(brandId)
}

// 保存品牌
func (b *BrandServiceImpl)SaveBrand(v *promodel.ProBrand) (int32, error) {
    id,err := b.rep.SaveProBrand(v)
    return int32(id),err
}

// 获取所有品牌
func (b *BrandServiceImpl)AllBrands() []*promodel.ProBrand {
    return b.rep.SelectProBrand("",)
}

// 获取关联的品牌编号
func (b *BrandServiceImpl)Brands(proModel int32) []*promodel.ProBrand {
    return b.rep.SelectProBrand("pro_model=?",proModel)
}

// 关联品牌
func (b *BrandServiceImpl)SetBrands(proModel int32, brandId []int32) error {
    return b.rep.SetModelBrands(proModel,brandId)
}
