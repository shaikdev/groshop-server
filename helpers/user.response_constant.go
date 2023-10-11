package helpers

const (
	// Success Response.
	SUCCESS                    = "Success"
	USER_CREATE_SUCCESS        = "User created successfully"
	USER_FETCH_SUCCESS         = "User fetched successfully"
	USER_LOGIN_SUCCESS         = "User login successfully"
	USER_GET_SUCCESSFULLY      = "Users get successfully"
	USER_EDIT_SUCCESSFULLY     = "Users edit successfully"
	USERS_DELETED_SUCCESSFULLY = "Users deleted successfully"
	USER_DELETED_SUCCESSFULLY  = "User deleted successfully"

	// Failed Response.
	FAILED                  = "Failed"
	USER_CREATE_FAILED      = "Failed to create an user"
	USER_FETCH_FAILED       = "User not found"
	PASSWORD_HASH_FAILED    = "Password hash failed"
	FAILED_TOKEN_CREATION   = "Failed to create token"
	EMAIL_ALREADY_EXIST     = "Email already exists"
	PASSWORD_DOES_NOT_MATCH = "Password does not match"
	NOT_AUTHORIZED          = "Not authorized"
	INVALID_TOKEN           = "Invalid token"
	TOKEN_NOT_FOUND         = "Token does not exist"
	USER_GET_FAILED         = "Failed to get user"
	USER_EDIT_FAILED        = "Users edit failed"
	USERS_DELETE_FAILED     = "Failed to delete users"
	USER_DELETE_FAILED      = "Failed to delete user"
)
