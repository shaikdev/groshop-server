package menuresponse

const (
	// SUCCESS
	CREATE_SUCCESS      = "Menu created successfully"
	GET_SUCCESS         = "Menu get succeeded"
	GET_MANY_SUCCESS    = "Menus fetch successfully"
	EDIT_SUCCESS        = "Menu edit succeeded"
	DELETE_SUCCESS      = "Menu deleted successfully"
	DELETE_MANY_SUCCESS = "Menus delete successfully"

	// FAILURE
	CREATE_FAILED      = "Menu created failed"
	GET_FAILED         = "Menu not found"
	GET_MANY_FAILED    = "Menus fetch failed"
	EDIT_FAILED        = "Menus edit failed"
	DELETE_FAILED      = "Menu delete failed"
	DELETE_MANY_FAILED = "Menus deleted failed"
	MENU_ALREADY_EXIST = "Menu name already exists"
)
