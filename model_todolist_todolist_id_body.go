/*
 * ToDoList Restful API
 *
 * OpenAPI for ToDoList Restful API
 *
 * API version: 1
 * Contact: mhdianrush@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TodolistTodolistIdBody struct {
	Name string `json:"name,omitempty"`
	Priority int32 `json:"priority,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
