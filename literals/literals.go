package literals

type routesliterals struct {
	Home                string
	Login               string
	Logout              string
	Register            string
	Todolists           string
	TodolistUpdate      string
	TodolistDelete      string
	Todolist            string
	TodolistID          string
	Createtodolist      string
	Updatetodolist      string
	UpdatetodolistID    string
	Deletetodolist      string
	DeletetodolistID    string
	Todolistitem        string
	TodolistitemID      string
	TodolistitemDone    string
	TodolistitemDoneID  string
	UpdateTodlistitemID string
	UpdateTodlistitem   string
}

var (
	SysRoutes routesliterals = routesliterals{
		Home:                "/",
		Login:               "/Login",
		Logout:              "/Logout",
		Register:            "/Register",
		Todolists:           "/TodoLists",
		Todolist:            "/Todolist",
		TodolistID:          "/TodoList/:Id",
		Createtodolist:      "/Create/TodoList",
		Updatetodolist:      "/Update/TodoList",
		UpdatetodolistID:    "/Update/TodoList/:Id",
		UpdateTodlistitemID: "Update/TodoList/TodoItem/:Id",
		UpdateTodlistitem:   "Update/TodoList/TodoItem",
		Deletetodolist:      "/Delete/TodoList",
		DeletetodolistID:    "/Delete/TodoList/:Id",
		Todolistitem:        "/TodoList/Item",
		TodolistitemID:      "/TodoList/Item/:Id",
		TodolistitemDone:    "/TodoList/Item/Done",
		TodolistitemDoneID:  "/TodoList/Item/Done/:Id",
	}
)
