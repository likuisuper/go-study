package controller

import (
	"basic/example/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Person struct {
	BaseController
}

// @Summary 添加
// @Tags person
// @Accept json
// @Produce json
// @Param person body model.Person true "人物添加req"
// @Success 200 {string} interface{}
// @Failure 500 {string} string
// @Router /v1/person/ [post]
func (p *Person) Add(g *gin.Context) {
	var data model.Person
	err := g.BindJSON(&data)
	if err != nil {
		p.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	if err := model.AddPerson(&data); err != nil {
		p.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	p.ResponseSuccess(g)
	return
}

// @Summary 删除
// @Router /v1/del/{id} [delete]
func (p *Person) Del(g *gin.Context) {
	pId := g.Param("id")
	if pId == "" {
		p.ResponseFailureForParameter(g, CErrParam)
		return
	}
	id, err := strconv.Atoi(pId)
	if err != nil {
		p.ResponseFailureForFuncErr(g, CErrTypeConversion)
		return
	}
	if err := model.DeletePerson(id); err != nil {
		p.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	p.ResponseSuccess(g)
	return
}

// @Summary 更新人物
// @Router /v1/person/{id} [put]
func (p *Person) Update(g *gin.Context) {
	sId := g.Param("id")
	if sId == "" {
		p.ResponseFailureForParameter(g, CErrParam)
		return
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		p.ResponseFailureForFuncErr(g, CErrTypeConversion)
		return
	}
	var data model.Person
	err = g.BindJSON(&data)
	if err != nil {
		p.ResponseFailureForParameter(g, CErrJSON)
		return
	}
	data.UserId = id
	if err := model.UpdatePerson(&data); err != nil {
		p.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	p.ResponseSuccess(g)
	return
}

// @Summary 通过id获取人物
// @Tags person
// @Param id path int true "Person UserId"
// @Success 200 {string} model.Person{}
// @Failure 500 {string} string
// @Router /v1/person/{id} [get]
func (p *Person) GetById(g *gin.Context) {
	sId := g.Param("id")
	if sId == "" {
		p.ResponseFailureForParameter(g, CErrParam)
		return
	}
	id, err := strconv.Atoi(sId)
	if err != nil {
		p.ResponseFailureForFuncErr(g, CErrTypeConversion)
		return
	}
	data, err := model.GetPersonById(id)
	if err != nil {
		p.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	p.ResponseData(g, data)
	return
}

// @Summary 获取人物
// @Tags person
// @Router /v1/person/ [get]
func (p *Person) GetAll(g *gin.Context) {
	data, err := model.GetPersonAll()
	if err != nil {
		p.ResponseFailureForParameter(g, err.Error())
		return
	}
	p.ResponseData(g, data)
	return
}
