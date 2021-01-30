package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

func validateRouterGetAds(ctx *gin.Context) (
	page int, perPage int, sortBy string, isReversed bool, errorData map[string]interface{},
) {
	var err error
	pageStr, ok := ctx.GetQuery("page")
	if !ok {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			errorData = getErrorResponseValidation("page", "positive int")
			return
		}
	}
	perPageStr, ok := ctx.GetQuery("per_page")
	if !ok {
		perPage = 10
	} else {
		perPage, err = strconv.Atoi(perPageStr)
		if err != nil || perPage <= 0 {
			errorData = getErrorResponseValidation("per_page", "positive int")
			return
		}
	}
	sortBy, ok = ctx.GetQuery("sort_by")
	if !ok {
		sortBy = "created_at"
	} else {
		if sortBy != "price" && sortBy != "created_at" {
			errorData = getErrorResponseValidation("sort_by", "\"price\" or \"created_at\"")
			return
		}
	}
	order, ok := ctx.GetQuery("order")
	if !ok {
		isReversed = true
	} else {
		if order == "asc" {
			isReversed = false
		} else if order == "desc" {
			isReversed = true
		} else {
			errorData = getErrorResponseValidation("order", "\"asc\" or \"desc\"")
			return
		}
	}
	return
}

func validateRouterGetAd(ctx *gin.Context) (
	adId uuid.UUID, isDescription bool, isURLs bool, errorData map[string]interface{},
) {
	adIdStr, ok := ctx.GetQuery("id")
	if !ok {
		errorData = getErrorResponseRequired("id")
		return
	}
	adId, err := uuid.Parse(adIdStr)
	if err != nil {
		errorData = getErrorResponseValidation("id", "UUID")
		return
	}
	fields := ctx.QueryArray("fields")
	for _, field := range fields {
		if field == "description" {
			isDescription = true
		} else if field == "urls" {
			isURLs = true
		} else {
			errorData = getErrorResponseValidation("fields", "equal \"description\" or \"urls\"")
			return
		}
	}
	return
}

func validateRouterCreateAd(ctx *gin.Context) (
	adName string, adDescription string, adPhotoUrls []string, adPrice int, errorData map[string]interface{},
) {
	adName, ok := ctx.GetPostForm("name")
	fmt.Println(adName)
	if !ok {
		errorData = getErrorResponseRequired("name")
		return
	}
	if len(adName) > 200 {
		errorData = getErrorResponseValidation("name", "no more than 200 symbols")
		return
	}
	adDescription, ok = ctx.GetPostForm("description")
	if !ok {
		errorData = getErrorResponseRequired("description")
		return
	}
	if len(adDescription) > 200 {
		errorData = getErrorResponseValidation("name", "no more than 1000 symbols")
		return
	}
	adPhotoUrls, ok = ctx.GetPostFormArray("photos_urls")
	if !ok {
		adPhotoUrls = []string{}
	} else {
		if len(adPhotoUrls) > 3 {
			errorData = getErrorResponseValidation("name", "no more than 3 urls")
			return
		}
	}
	adPriceStr, ok := ctx.GetPostForm("price")
	if !ok {
		errorData = getErrorResponseRequired("price")
		return
	}
	adPrice, err := strconv.Atoi(adPriceStr)
	if err != nil {
		errorData = getErrorResponseValidation("price", "int")
		return
	}
	return
}
