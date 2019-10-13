package main

import "net/http"

//Route Used to match requets with approaite handlers
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes slice holds all information on routes and what handler function to use
type Routes []Route

var routes = Routes{

	Route{
		"StudentCreation",
		"POST",
		"/students/create",
		StudentCreation,
	},
	Route{
		"TeacherCreation",
		"POST",
		"/teachers/create",
		TeacherCreation,
	},
	Route{
		"ClassCreation",
		"POST",
		"/classes/create",
		ClassCreation,
	},
	Route{
		"StudentSearch",
		"GET",
		"/students/search",
		StudentSearch,
	},
	Route{
		"TeacherSearch",
		"GET",
		"/teachers/search",
		TeacherSearch,
	},
	Route{
		"ClassSearch",
		"GET",
		"/classes/search",
		ClassSearch,
	},
	Route{
		"StudentUpdate",
		"PUT",
		"/students/update",
		StudentUpdate,
	},
}
