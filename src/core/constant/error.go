package constant

const (
	ErrUserNotFound    = "user not found"
	ErrUnexpected      = "unexpected error"
	ErrExistedEmail    = "email already existed"
	ErrCacheKeyNil     = "cache key is nil"
	ErrUserVerified    = "user is verified"
	ErrOtpNotFound     = "otp not found"
	ErrOtpInvalid      = "otp invalid"
	ErrInvalidPassword = "invalid password"

	ErrWrongPassword = "wrong password"
	ErrGroupNotFound = "group not found"

	ErrGroupMemberNotFound = "group member not found"

	ErrForbiddenAddMember = "forbidden add member"

	ErrForbiddenRemoveMember       = "forbidden remove member"
	ErrForbiddenGetMember          = "forbidden get member"
	ErrForbiddenCreateShoppingList = "forbidden create shopping list"

	ErrShoppingListNotFound        = "shopping list not found"
	ErrForbiddenUpdateShoppingList = "forbidden update shopping list"
	ErrForbiddenDeleteShoppingList = "forbidden delete shopping list"

	ErrForbiddenCreateShoppingTask = "forbidden create shopping task"
	ErrForbiddenGetShoppingTask    = "forbidden get shopping task"
	ErrForbiddenGetShoppingList    = "forbidden get shopping list"
)
