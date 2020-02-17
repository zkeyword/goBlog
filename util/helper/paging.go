package helper

import (
	"math"
)

// Paging 分页类型
type Paging struct {
	Page      int64   //当前页
	Pagesize  int64   //每页条数
	Total     int64   //总条数
	PageCount int64   //总页数
	Num       []int64 //分页序数
	NumCount  int64   //总页序数
}

// CreatePaging 创建分页
func CreatePaging(page, pagesize, total int64) *Paging {
	if page < 1 {
		page = 1
	}
	if pagesize < 1 {
		pagesize = 10
	}

	pageCount := math.Ceil(float64(total) / float64(pagesize))

	paging := new(Paging)
	paging.Page = page
	paging.Pagesize = pagesize
	paging.Total = total
	paging.PageCount = int64(pageCount)
	paging.NumCount = 7
	paging.setNum()
	return paging
}

// setNum 设置
func (p *Paging) setNum() {
	p.Num = []int64{}
	if p.PageCount == 0 {
		return
	}

	half := math.Floor(float64(p.NumCount) / float64(2))
	begin := p.Page - int64(half)
	if begin < 1 {
		begin = 1
	}

	end := begin + p.NumCount - 1
	if end >= p.PageCount {
		begin = p.PageCount - p.NumCount + 1
		if begin < 1 {
			begin = 1
		}
		end = p.PageCount
	}

	for i := begin; i <= end; i++ {
		p.Num = append(p.Num, i)
	}
}
