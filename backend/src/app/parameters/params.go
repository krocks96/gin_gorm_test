package parameters

type PostUserParams struct {
	DisplayName string `form:"display_name" binding:"required"`
}
