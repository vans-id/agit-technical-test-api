package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vans-id/agit-technical-test-api.git/internal/dto"
	"github.com/vans-id/agit-technical-test-api.git/internal/entity"
	"github.com/vans-id/agit-technical-test-api.git/internal/usecase"
	"github.com/vans-id/agit-technical-test-api.git/shared/apperror"
)

type EmplHandler interface {
	HandleGetAll(ctx *gin.Context)
	HandleGetDetail(ctx *gin.Context)
	HandleAdd(ctx *gin.Context)
	HandleEdit(ctx *gin.Context)
	HandleRemove(ctx *gin.Context)
}

type emplHandler struct {
	emplUsecase usecase.EmplUsecase
}

func NewEmplHandler(uc usecase.EmplUsecase) EmplHandler {
	return &emplHandler{emplUsecase: uc}
}

func (h *emplHandler) HandleGetAll(ctx *gin.Context) {
	search := ctx.Query("search")
	employees, err := h.emplUsecase.GetAll(ctx.Request.Context(), search)
	if err != nil {
		ctx.Error(err)
		return
	}
	responses := []*dto.EmployeeResponse{}
	for _, entity := range employees {
		responses = append(responses, entity.ToResponse())
	}
	resp := dto.JSONResponse{
		Data: responses,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *emplHandler) HandleGetDetail(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.Error(err)
		return
	}
	employee, err := h.emplUsecase.GetDetail(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.JSONResponse{
		Data: employee.ToResponse(),
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *emplHandler) HandleAdd(ctx *gin.Context) {
	param := dto.EmployeeRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	payload := &entity.Employee{
		Name:         param.Name,
		Nip:          param.Nip,
		PlaceOfBirth: param.PlaceOfBirth,
		DateOfBirth:  param.DateOfBirth,
		Address:      param.Address,
		Religion:     param.Religion,
		Gender:       param.Gender,
		Phone:        param.Phone,
		Email:        param.Email,
	}
	err = h.emplUsecase.Add(ctx.Request.Context(), payload)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.JSONResponse{
		Data: payload.ToResponse(),
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (h *emplHandler) HandleEdit(ctx *gin.Context) {
	param := dto.EmployeeRequest{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.Error(err)
		return
	}
	payload := &entity.Employee{
		Id:           uint(id),
		Name:         param.Name,
		Nip:          param.Nip,
		PlaceOfBirth: param.PlaceOfBirth,
		DateOfBirth:  param.DateOfBirth,
		Address:      param.Address,
		Religion:     param.Religion,
		Gender:       param.Gender,
		Phone:        param.Phone,
		Email:        param.Email,
	}
	err = h.emplUsecase.Edit(ctx.Request.Context(), payload)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.JSONResponse{
		Data: payload.ToResponse(),
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *emplHandler) HandleRemove(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.Error(err)
		return
	}
	err = h.emplUsecase.Remove(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.JSONResponse{
		Data: gin.H{"id": id},
	}
	ctx.JSON(http.StatusOK, resp)
}
