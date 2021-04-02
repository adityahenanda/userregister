package http_handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"userregister/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func (h *HandlerUserAccount) RegisterUserAccountHandler(c *gin.Context) {
	var response models.ResponseResult
	var req models.UserAccountRequest
	body, _ := ioutil.ReadAll(c.Request.Body)

	//get body request
	err := json.Unmarshal(body, &req)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		response.Result = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// validate tag required
	v := validator.New()
	err = v.Struct(req)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		response.Result = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	id, err := h.UserAccountUseCase.RegisterUserAccount(&req)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		response.Result = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "success"
	response.Result = id
	c.JSON(http.StatusOK, response)
}

func (h *HandlerUserAccount) LoginUserAccountHandler(c *gin.Context) {
	var response models.ResponseResult
	var req models.UserAccountLoginRequest
	body, _ := ioutil.ReadAll(c.Request.Body)

	//get body request
	err := json.Unmarshal(body, &req)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		response.Result = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	data, err := h.UserAccountUseCase.LoginUserAccount(&req)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		response.Result = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//generate token
	token := h.JwtService.GenerateToken(strconv.FormatUint(data.Id, 10))

	loginRes := make(map[string]interface{})
	loginRes["token"] = token
	loginRes["data"] = data
	response.Status = http.StatusOK
	response.Message = "success"
	response.Result = loginRes
	c.JSON(http.StatusOK, response)
}

func (h *HandlerUserAccount) GetUserAccountById(c *gin.Context) {

	var response models.ResponseResult

	paramID := c.Param("userId")
	//parse params id into int using uint to check value whether string or negative number
	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	res, err := h.UserAccountUseCase.GetUserAccountById(id)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "success"
	response.Result = res
	c.JSON(http.StatusOK, response)
}

func (h *HandlerUserAccount) GetAllUserAccount(c *gin.Context) {
	var response models.AllUserAccountResp

	//prevent sql injection with queryEscape
	pageParam := url.QueryEscape(c.DefaultQuery("page", "1"))
	limitParam := url.QueryEscape(c.DefaultQuery("limit", "100"))

	//parse page and limit into int using uint to check value whether string or negative number
	page, err := strconv.ParseUint(pageParam, 10, 64)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	limit, err := strconv.ParseUint(limitParam, 10, 64)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	res, total, err := h.UserAccountUseCase.GetAllUserAccount(int(limit), int(page))
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "success"
	response.Result = res
	response.TotalData = total
	c.JSON(http.StatusOK, response)
	return

}

func (h *HandlerUserAccount) DeleteUserAccount(c *gin.Context) {
	var response models.ResponseResult

	paramID := c.Param("userId")
	//parse params id into int using uint to check value whether string or negative number
	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	err = h.UserAccountUseCase.DeleteUserAccount(id)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "success"
	response.Result = id
	c.JSON(http.StatusOK, response)
}

func (h *HandlerUserAccount) UpdateUserAccount(c *gin.Context) {
	var response models.ResponseResult
	var req models.UserAccountUpdateRequest
	body, _ := ioutil.ReadAll(c.Request.Body)
	//get body request
	err := json.Unmarshal(body, &req)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	paramID := c.Param("userId")
	//parse params id into int using uint to check value whether string or negative number
	id, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//begin bussiness logic
	err = h.UserAccountUseCase.UpdateUserAccount(id, &req)
	if err != nil {
		response.Message = err.Error()
		response.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "success"
	response.Result = id
	c.JSON(http.StatusOK, response)
	return
}
