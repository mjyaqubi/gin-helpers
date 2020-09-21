package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BadRequest (400)
type BadRequest struct {
	Status int    `json:"status" example:"400"`
	Reason string `json:"reason" example:"bad request"`
}

// Unauthorized (401)
type Unauthorized struct {
	Status int    `json:"status" example:"401"`
	Reason string `json:"reason" example:"unauthorized request"`
}

// Forbidden (403)
type Forbidden struct {
	Status int    `json:"status" example:"403"`
	Reason string `json:"reason" example:"forbidden request"`
}

// NotFound (404)
type NotFound struct {
	Status int    `json:"status" example:"404"`
	Reason string `json:"reason" example:"not found"`
}

// NotAcceptable (406)
type NotAcceptable struct {
	Status int    `json:"status" example:"406"`
	Reason string `json:"reason" example:"not acceptable request"`
}

// InternalServerError (500)
type InternalServerError struct {
	Status int    `json:"status" example:"500"`
	Reason string `json:"reason" example:"internal server error"`
}

// Response type struct
type Response struct{}

// New - new response helper
func New() *Response {
	return &Response{}
}

func send(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func abortWithError(c *gin.Context, status int, reason string) {
	c.AbortWithStatusJSON(status, &gin.H{"status": status, "reason": reason})
}

// Ok response 200
func (res *Response) Ok(c *gin.Context, data interface{}) {
	send(c, http.StatusOK, data)
}

// Created response 201
func (res *Response) Created(c *gin.Context, data interface{}) {
	send(c, http.StatusCreated, data)
}

// BadRequest response 400
func (res *Response) BadRequest(c *gin.Context, reason string) {
	if reason == "" {
		reason = "bad request"
	}
	abortWithError(c, http.StatusBadRequest, reason)
}

// Unauthorized response 401
func (res *Response) Unauthorized(c *gin.Context, reason string) {
	if reason == "" {
		reason = "unauthorized request"
	}
	abortWithError(c, http.StatusUnauthorized, reason)
}

// Forbidden response 403
func (res *Response) Forbidden(c *gin.Context, reason string) {
	if reason == "" {
		reason = "forbidden request"
	}
	abortWithError(c, http.StatusForbidden, reason)
}

// NotFound response 303
func (res *Response) NotFound(c *gin.Context, reason string) {
	if reason == "" {
		reason = "not found"
	}
	abortWithError(c, http.StatusNotFound, reason)
}

// InternalServerError response 500
func (res *Response) InternalServerError(c *gin.Context, reason string) {
	if reason == "" {
		reason = "internal server error"
	}
	abortWithError(c, http.StatusInternalServerError, reason)
}
