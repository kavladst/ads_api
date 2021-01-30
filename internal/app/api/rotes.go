package api

import (
	"github.com/gin-gonic/gin"
)

func (a *Api) initRouter() {
	router := gin.Default()
	v1 := router.Group("/v1")
	ads := v1.Group("/ads")
	{
		ads.GET("/", a.getAds)
		ads.GET("/detail", a.getAd)
		ads.POST("/", a.createAd)
	}
	a.router = router
}

func (a *Api) getAds(ctx *gin.Context) {
	page, perPage, sortBy, isReversed, errorData := validateRouterGetAds(ctx)
	if errorData != nil {
		ctx.JSON(400, errorData)
		return
	}

	ads, err := a.Storage.GetAds(page, perPage, sortBy, isReversed)
	if err != nil {
		ctx.JSON(400, getResponseData(nil, err))
		return
	}
	responseData := make([]map[string]interface{}, len(ads))
	for i, ad := range ads {
		imageURL := "ad has not image"
		if len(ad.PhotosUrls) > 0 {
			imageURL = ad.PhotosUrls[0]
		}
		responseData[i] = map[string]interface{}{"id": ad.ID, "name": ad.Name, "image_url": imageURL, "price": ad.Price}
	}
	ctx.JSON(200, getResponseData(responseData, nil))
	return
}

func (a *Api) getAd(ctx *gin.Context) {
	adId, isDescription, isURLs, errorData := validateRouterGetAd(ctx)
	if errorData != nil {
		ctx.JSON(400, errorData)
		return
	}

	ad, err := a.Storage.GetAd(adId)
	if err != nil {
		ctx.JSON(400, getResponseData(nil, err))
		return
	}
	responseData := map[string]interface{}{
		"name":  ad.Name,
		"price": ad.Price,
	}
	if isDescription {
		responseData["description"] = ad.Description
	}
	if isURLs {
		responseData["urls"] = ad.PhotosUrls
	} else {
		if len(ad.PhotosUrls) > 0 {
			responseData["url"] = ad.PhotosUrls[0]
		} else {
			responseData["url"] = "ad has not image"
		}
	}
	ctx.JSON(200, getResponseData(responseData, nil))
	return
}

func (a *Api) createAd(ctx *gin.Context) {
	adName, adDescription, adPhotoUrls, adPrice, errorData := validateRouterCreateAd(ctx)
	if errorData != nil {
		ctx.JSON(400, errorData)
		return
	}

	adId, err := a.Storage.CreateAd(adName, adDescription, adPhotoUrls, adPrice)
	if err != nil {
		ctx.JSON(400, getResponseData(nil, err))
		return
	}
	ctx.JSON(201, getResponseData(map[string]interface{}{"new_ad_id": adId.String(), "code": 201}, nil))
	return
}
