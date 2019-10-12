package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ClassIndex",
		"GET",
		"/classes",
		ClassIndex,
	},
	Route{
		"StudentCreation",
		"POST",
		"/students/create",
		StudentCreation,
	},
	Route{
		"StudentIndex",
		"GET",
		"/students",
		StudentIndex,
	},
	Route{
		"TeacherIndex",
		"GET",
		"/teachers",
		TeacherIndex,
	},
	Route{
		"test",
		"GET",
		"/test",
		test,
	},
}
