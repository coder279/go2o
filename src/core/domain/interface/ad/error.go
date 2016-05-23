/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-12 16:53
 * description :
 * history :
 */

package ad

import (
	"go2o/src/core/infrastructure/domain"
)

var (
	ErrNameExists *domain.DomainError = domain.NewDomainError(
		"name_exists", "已经存在相同的名称的广告")

	ErrInternalDisallow *domain.DomainError = domain.NewDomainError(
		"err_internal_disallow", "不允许修改系统内置广告")

	ErrNoSuchAd *domain.DomainError = domain.NewDomainError(
		"err_no_such_ad", "广告不存在")

ErrNoSuchAdGroup *domain.DomainError = domain.NewDomainError(
	"err_no_such_ad_group", "广告组不存在")

	ErrNoSuchAdPosition *domain.DomainError = domain.NewDomainError(
		"err_no_such_ad_position", "广告位不存在")
)
