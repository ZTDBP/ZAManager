// Copyright 2022-present The ZTDBP Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mdb

type Paginate struct {
	Page     int `json:"page" form:"page,default=1"`            // 分页号
	LimitNum int `json:"limit_num" form:"limit_num,default=10"` // 每页限制数量
	Offset   int // 计算获取offset
}

func (p *Paginate) Init() {
	if p.LimitNum == 0 {
		p.LimitNum = 10 // 默认一页10条
	}
	if p.Page == 0 {
		p.Page = 1
	}
	p.Offset = p.getOffset()
}

func (p *Paginate) getOffset() int {
	offset := (p.Page - 1) * p.LimitNum
	if offset < 0 {
		offset = 0
	}
	return offset
}

func (p *Paginate) GetOffset() int {
	offset := (p.Page - 1) * p.LimitNum
	if offset < 0 {
		offset = 0
	}
	p.Offset = offset
	return offset
}

func (p *Paginate) GetPageNum(limitNum int, defaultNum int) int {
	if limitNum > 20 {
		return 10
	} else if limitNum <= 0 {
		return defaultNum
	} else {
		return limitNum
	}
}

func (p *Paginate) DealOffset() {
	if p.Page == 0 {
		p.Page = 1
	}
	p.LimitNum = p.GetPageNum(p.LimitNum, 10)
	p.GetOffset()
	return
}
