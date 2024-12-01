package response

import "github.com/KhaiHust/authen_service/core/entity"

type GroupMemberResponse struct {
	ID      int64                `json:"id"`
	GroupID int64                `json:"group_id"`
	UserID  int64                `json:"user_id"`
	Role    *GroupRoleResponse   `json:"role"`
	User    *UserProfileResponse `json:"user_profile"`
}

type GroupRoleResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type UserProfileResponse struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"user_id"`
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	AvatarImageUrl string `json:"avatar_image_url,omitempty"`
}

func ToGroupMemberResponse(entity *entity.GroupMemberEntity) *GroupMemberResponse {
	return &GroupMemberResponse{
		ID:      entity.ID,
		GroupID: entity.GroupID,
		UserID:  entity.UserID,
		Role:    ToGroupRoleResponse(entity.Role),
		User:    ToUserProfileResponse(entity.UserProfile),
	}
}

func ToUserProfileResponse(profile *entity.UserProfileEntity) *UserProfileResponse {
	return &UserProfileResponse{
		ID:             profile.ID,
		UserID:         profile.UserID,
		Name:           profile.Name,
		Email:          profile.Email,
		AvatarImageUrl: profile.AvatarImageUrl,
	}
}

func ToGroupRoleResponse(role *entity.GroupRoleEntity) *GroupRoleResponse {
	return &GroupRoleResponse{
		ID:   role.ID,
		Name: role.Name,
		Code: role.Code,
	}
}
func ToListGroupMemberResponse(entities []*entity.GroupMemberEntity) []*GroupMemberResponse {
	responses := make([]*GroupMemberResponse, 0)
	for _, group := range entities {
		responses = append(responses, ToGroupMemberResponse(group))
	}
	return responses
}
