package mysql

import (
	"errors"

	"github.com/ztalab/ZAManager/app/v1/access/model/mmysql"
	"github.com/ztalab/ZAManager/app/v1/access/model/mparam"
	"github.com/ztalab/ZAManager/pkg/logger"
	"github.com/ztalab/ZAManager/pkg/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Client struct {
	c *gin.Context
	mysql.DaoMysql
}

func NewClient(c *gin.Context) *Client {
	return &Client{
		DaoMysql: mysql.DaoMysql{TableName: "zta_client"},
		c:        c,
	}
}

func (p *Client) GetClient(param mparam.GetClient) (
	total int64, list []mmysql.Client, err error) {
	orm := p.GetOrm().DB
	query := orm.Table(p.TableName)
	if len(param.Name) > 0 {
		query = query.Where("name like ?", "%"+param.Name+"%")
	}
	err = query.Model(&list).Count(&total).Error
	if total > 0 {
		offset := param.GetOffset()
		err = query.Limit(param.LimitNum).Offset(offset).
			Order("created_at desc").
			Find(&list).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetClient err : %v", err)
	}
	return
}

func (p *Client) GetClientByID(id uint64) (info mmysql.Client, err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Where("id = ?", id).First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetClientById err : %v", err)
	}
	return
}

func (p *Client) AddClient(data *mmysql.Client) (err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Create(&data).Error
	if err != nil {
		logger.Errorf(p.c, "AddClient err : %v", err)
	}
	return
}

func (p *Client) EditClient(data mmysql.Client) (err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Save(&data).Error
	if err != nil {
		logger.Errorf(p.c, "EditClient err : %v", err)
	}
	return
}

func (p *Client) DelClient(id uint64) (err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Where("id = ?", id).Delete(&mmysql.Client{}).Error
	if err != nil {
		logger.Errorf(p.c, "DelClient err : %v", err)
	}
	return
}
