package controller

import (
	"errors"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/KhaiHust/authen_service/public/resource/request"
	"github.com/KhaiHust/authen_service/public/resource/response"
	"github.com/KhaiHust/authen_service/public/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golibs-starter/golib/log"
	"strconv"
)

type GroupController struct {
	*BaseController
	groupService service.IGroupService
}

func NewGroupController(base *BaseController, groupService service.IGroupService) *GroupController {
	return &GroupController{groupService: groupService,
		BaseController: base}
}
func (g GroupController) CreateGroup(c *gin.Context) {
	var req request.CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "error when binding request", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := g.GetUserID(c)
	if err != nil {
		log.Error(c, "error when getting user id", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}

	group, err := g.groupService.CreateGroup(c, userID, request.ToCreateGroupRequestDto(&req))
	if err != nil {
		log.Error(c, "error when creating group", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToCreateGroupResponse(group))
}
func (g GroupController) GetUserID(c *gin.Context) (int64, error) {
	claims, ok := c.Get("claims")
	if !ok {
		log.Error(c, "error when getting claims from context")
		return 0, errors.New("error when getting claims from context")
	}
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		log.Error(c, "error when casting claims to map")

		return 0, errors.New("error when casting claims to map")
	}
	userID, ok := claimsMap[constant.CLAIM_USER_ID]
	if !ok {
		log.Error(c, "error when getting user id from claims")
		return 0, errors.New("error when getting user id from claims")
	}

	return int64(userID.(float64)), nil
}

func (g GroupController) AddMember(c *gin.Context) {
	var req request.AddMemberGroupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "error when binding request", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := g.Validator.Struct(req); err != nil {
		log.Error(c, "error when validating request", err)
		errCode, err := strconv.ParseInt(err.Error(), 10, 64)
		if err != nil {
			log.Error(c, "error when parsing error code", err)
			apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
			return
		}
		apihelper.AbortErrorHandle(c, int(errCode))
		return
	}
	userID, err := g.GetUserID(c)
	if err != nil {
		log.Error(c, "error when getting user id", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	groupMember, err := g.groupService.AddNewMember(c, userID, req.GroupID, req.Email)
	if err != nil {
		log.Error(c, "error when adding member", err)
		if err.Error() == constant.ErrForbiddenAddMember {
			apihelper.AbortErrorHandle(c, common.GeneralForbidden)
			return
		}
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToAddNewMemberResp(groupMember))
}

func (g GroupController) GetListMember(c *gin.Context) {
	groupID, err := strconv.ParseInt(c.Param("groupID"), 10, 64)
	if err != nil {
		log.Error(c, "error when parsing group id", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := g.GetUserID(c)
	if err != nil {
		log.Error(c, "error when getting user id", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	groupMembers, err := g.groupService.GetMembers(c, userID, groupID)
	if err != nil {
		log.Error(c, "error when getting list member", err)
		if err.Error() == constant.ErrForbiddenGetMember {
			apihelper.AbortErrorHandle(c, common.GeneralForbidden)
			return
		}
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToListGroupMemberResponse(groupMembers))
}
