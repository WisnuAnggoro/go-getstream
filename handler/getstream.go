package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/wisnuanggoro/go-getstream/getstream"

	"github.com/gin-gonic/gin"
)

type handler struct {
	getstreamSvc getstream.Service
}

type GetstreamHandler interface {
	AddPostByUserSerial(c *gin.Context)
	GetPostByUserSerial(c *gin.Context)
	GetPostDetailByUserSerial(c *gin.Context)
	DeletePostByPostID(c *gin.Context)
	GetTimelineByUserSerial(c *gin.Context)
	GetDetailTimelineByUserSerial(c *gin.Context)
	Follow(c *gin.Context)
	Unfollow(c *gin.Context)
	GetFeedFollowersByUserSerial(c *gin.Context)
	GetFollowedFeedsByUserSerial(c *gin.Context)
	AddLikeToPostID(c *gin.Context)
	RetrieveLikeDetailOnPostID(c *gin.Context)
	RetrieveLikeDetailOnPostIDWithPagination(c *gin.Context)
	RemoveLikeByReactionID(c *gin.Context)
}

func NewGetstreamHandler(getstreamSvc getstream.Service) GetstreamHandler {
	return &handler{
		getstreamSvc: getstreamSvc,
	}
}

func (h *handler) AddPostByUserSerial(c *gin.Context) {
	userSerial := c.Query("userSerial")
	postContent := c.Query("postContent")
	if userSerial == "" || postContent == "" {
		AddResponseToContext(c, http.StatusBadRequest, "userSerial and postContent are mandatory", nil)
		return
	}

	resp, err := h.getstreamSvc.AddPostByUserSerial(userSerial, postContent)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusCreated, "Feed has been successfully added to timeline!", resp)
}

func (h *handler) GetPostByUserSerial(c *gin.Context) {
	userSerial := c.Param("userSerial")
	if userSerial == "" {
		AddResponseToContext(c, http.StatusBadRequest, "userSerial is mandatory", nil)
		return
	}

	resp, err := h.getstreamSvc.GetPostByUserSerial(userSerial)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, "success", resp)
}

func (h *handler) GetPostDetailByUserSerial(c *gin.Context) {
	userSerial := c.Param("userSerial")
	if userSerial == "" {
		AddResponseToContext(c, http.StatusBadRequest, "userSerial is mandatory", nil)
		return
	}

	resp, err := h.getstreamSvc.GetPostDetailByUserSerial(userSerial)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, "success", resp)
}

func (h *handler) DeletePostByPostID(c *gin.Context) {
	userSerial := c.Query("userSerial")
	postID := c.Query("postID")
	if userSerial == "" || postID == "" {
		AddResponseToContext(c, http.StatusBadRequest, "userSerial and postID are mandatory", nil)
		return
	}

	err := h.getstreamSvc.DeletePostByPostID(userSerial, postID)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, fmt.Sprintf("Post with ID %s has been successfully deleted!", postID), nil)
}

func (h *handler) GetTimelineByUserSerial(c *gin.Context) {
	userSerial := c.Param("userSerial")
	if userSerial == "" {
		AddResponseToContext(c, http.StatusBadRequest, "userSerial is are mandatory", nil)
		return
	}

	resp, err := h.getstreamSvc.GetTimelineByUserSerial(userSerial)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, "Success", resp)
}

func (h *handler) GetDetailTimelineByUserSerial(c *gin.Context) {
	userSerial := c.Param("userSerial")
	if userSerial == "" {
		AddResponseToContext(c, http.StatusBadRequest, "userSerial is are mandatory", nil)
		return
	}

	resp, err := h.getstreamSvc.GetDetailTimelineByUserSerial(userSerial)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, "Success", resp)
}

func (h *handler) Follow(c *gin.Context) {
	ownUserSerial := c.Query("ownUserSerial")
	targetUserSerial := c.Query("targetUserSerial")
	if ownUserSerial == "" || targetUserSerial == "" {
		AddResponseToContext(c, http.StatusBadRequest, "ownUserSerial and targetUserSerial are mandatory", nil)
		return
	}

	err := h.getstreamSvc.Follow(ownUserSerial, targetUserSerial)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, fmt.Sprintf("%s has successfully followed %s!", ownUserSerial, targetUserSerial), nil)
}

func (h *handler) Unfollow(c *gin.Context) {
	ownUserSerial := c.Query("ownUserSerial")
	targetUserSerial := c.Query("targetUserSerial")
	if ownUserSerial == "" || targetUserSerial == "" {
		AddResponseToContext(c, http.StatusBadRequest, "ownUserSerial and targetUserSerial are mandatory", nil)
		return
	}

	err := h.getstreamSvc.Unfollow(ownUserSerial, targetUserSerial)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, fmt.Sprintf("%s has successfully unfollowed %s!", ownUserSerial, targetUserSerial), nil)
}

func (h *handler) GetFeedFollowersByUserSerial(c *gin.Context) {
	userSerial := c.Param("userSerial")
	if userSerial == "" {
		AddResponseToContext(c, http.StatusBadRequest, "userSerial is are mandatory", nil)
		return
	}

	resp, err := h.getstreamSvc.GetFeedFollowersByUserSerial(userSerial)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, "Success", resp)
}

func (h *handler) GetFollowedFeedsByUserSerial(c *gin.Context) {
	userSerial := c.Param("userSerial")
	if userSerial == "" {
		AddResponseToContext(c, http.StatusBadRequest, "userSerial is are mandatory", nil)
		return
	}

	resp, err := h.getstreamSvc.GetFollowedFeedsByUserSerial(userSerial)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, "Success", resp)
}

func (h *handler) AddLikeToPostID(c *gin.Context) {
	likerUserSerial := c.Query("likerUserSerial")
	postID := c.Query("postID")
	if likerUserSerial == "" || postID == "" {
		AddResponseToContext(c, http.StatusBadRequest, "likerUserSerial and postID are mandatory", nil)
		return
	}

	resp, err := h.getstreamSvc.AddLikeToPostID(likerUserSerial, postID)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, fmt.Sprintf("%s has been successfully liked by %s!", postID, likerUserSerial), resp)
}

func (h *handler) RetrieveLikeDetailOnPostID(c *gin.Context) {
	postID := c.Param("postID")
	pageSizeString := c.Query("pageSize")

	if postID == "" {
		AddResponseToContext(c, http.StatusBadRequest, "postID is mandatory", nil)
		return
	}

	pageSize := 10
	if pageSizeString != "" {
		i, err := strconv.Atoi(pageSizeString)
		if err == nil && i > 0 {
			pageSize = i
		}
	}

	resp, err := h.getstreamSvc.RetrieveLikeDetailOnPostID(postID, pageSize)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, "Success", resp)
}

func (h *handler) RetrieveLikeDetailOnPostIDWithPagination(c *gin.Context) {
	postID := c.Param("postID")
	nextLikeID := c.Param("nextLikeID")
	pageSizeString := c.Query("pageSize")

	if postID == "" {
		AddResponseToContext(c, http.StatusBadRequest, "postID is mandatory", nil)
		return
	}

	pageSize := 10
	if pageSizeString != "" {
		i, err := strconv.Atoi(pageSizeString)
		if err == nil && i > 0 {
			pageSize = i
		}
	}

	resp, err := h.getstreamSvc.RetrieveLikeDetailOnPostIDWithPagination(postID, nextLikeID, pageSize)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, "Success", resp)
}

func (h *handler) RemoveLikeByReactionID(c *gin.Context) {
	reactionID := c.Param("reactionID")
	if reactionID == "" {
		AddResponseToContext(c, http.StatusBadRequest, "reactionID is mandatory", nil)
		return
	}

	err := h.getstreamSvc.RemoveLikeByReactionID(reactionID)
	if err != nil {
		AddResponseToContext(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	AddResponseToContext(c, http.StatusOK, fmt.Sprintf("reactionID %s has been successfully removed!", reactionID), nil)
}
