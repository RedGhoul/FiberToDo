package literals

type routesliterals struct {
	Home             string
	Login            string
	Logout           string
	Register         string
	Todolists        string
	Todolist         string
	Createtodolists  string
	Todolistid       string
	Updatetodolist   string
	Deletetodolist   string
	Todolistitem     string
	Todolistitemdone string
}

var (
	SysRoutes routesliterals = routesliterals{
		"/",
		"/Login",
		"/Logout",
		"/Register",
		"/TodoLists",
		"/Todolist",
		"/Create/ToDoList",
		"/TodoList/:Id",
		"/Update/TodoList/:Id",
		"/Delete/TodoList/:Id",
		"/TodoList/Item/:Id",
		"/TodoList/Item/Done/:Id",
	}
)
