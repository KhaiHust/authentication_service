package response

import "github.com/KhaiHust/authen_service/core/entity"

// AddNewMemberResp struct
type AddNewMemberResp struct {
	GroupMemberID int64 `json:"group_member_id"`
}

func ToAddNewMemberResp(groupMember *entity.GroupMemberEntity) *AddNewMemberResp {
	return &AddNewMemberResp{GroupMemberID: groupMember.ID}
}
