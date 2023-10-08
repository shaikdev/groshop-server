package helpers

const (
	// Success Response.
	SUCCESS             = "Success"
	USER_CREATE_SUCCESS = "user created successfully"
	USER_FETCH_SUCCESS  = "user fetched successfully"
	USER_LOGIN_SUCCESS  = "user login successfully"

	// Failed Response.
	FAILED                = "Failed"
	USER_CREATE_FAILED    = "Failed to create an user"
	USER_FETCH_FAILED     = "Failed to fetch user"
	PASSWORD_HASH_FAILED  = "password hash failed"
	FAILED_TOKEN_CREATION = "failed to create token"
	EMAIL_ALREADY_EXIST   = "email already exists"
)
