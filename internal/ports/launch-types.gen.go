// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package ports

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for KindStatus.
const (
	KindStatusDraft     KindStatus = "draft"
	KindStatusPublished KindStatus = "published"
)

// Defines values for LaunchStatus.
const (
	Completed LaunchStatus = "completed"
	Started   LaunchStatus = "started"
	Todo      LaunchStatus = "todo"
)

// Defines values for OrderStatus.
const (
	OrderStatusDraft     OrderStatus = "draft"
	OrderStatusPublished OrderStatus = "published"
)

// Kind model.
type Kind struct {
	// Description for the given kind.
	Description *string `json:"description,omitempty"`

	// Unique identifier for the given kind.
	Id openapi_types.UUID `json:"id"`

	// Unique name for the given kind.
	Name   string     `json:"name"`
	Status KindStatus `json:"status"`
}

// KindStatus defines model for Kind.Status.
type KindStatus string

// Launch model.
type Launch struct {
	// Description for the given launch.
	Description *string `json:"description,omitempty"`

	// Unique identifier for the given launch.
	Id openapi_types.UUID `json:"id"`

	// Unique name for the given launch.
	Name   string       `json:"name"`
	Status LaunchStatus `json:"status"`
}

// LaunchStatus defines model for Launch.Status.
type LaunchStatus string

// Order model.
type Order struct {
	// Description for the given order.
	Description *string `json:"description,omitempty"`

	// Unique identifier for the given order.
	Id openapi_types.UUID `json:"id"`

	// Unique name for the given order.
	Name   string      `json:"name"`
	Status OrderStatus `json:"status"`
}

// OrderStatus defines model for Order.Status.
type OrderStatus string

// Kind model.
type GetKindResponse = Kind

// Launch model.
type GetLaunchResponse = Launch

// Order model.
type GetOrderResponse = Order

// ListKindsResponse defines model for ListKindsResponse.
type ListKindsResponse struct {
	Kinds      []Kind `json:"kinds"`
	NextCursor string `json:"nextCursor"`
}

// ListLaunchesResponse defines model for ListLaunchesResponse.
type ListLaunchesResponse struct {
	Launchs    *[]Launch `json:"launchs,omitempty"`
	NextCursor string    `json:"nextCursor"`
}

// ListOrdersResponse defines model for ListOrdersResponse.
type ListOrdersResponse struct {
	NextCursor string  `json:"nextCursor"`
	Orders     []Order `json:"orders"`
}

// ChangeKindDescription defines model for ChangeKindDescription.
type ChangeKindDescription struct {
	Description string `json:"description"`
}

// ChangeKindName defines model for ChangeKindName.
type ChangeKindName struct {
	Name string `json:"name"`
}

// ChangeLaunchDescription defines model for ChangeLaunchDescription.
type ChangeLaunchDescription struct {
	Description string `json:"description"`
}

// ChangeLaunchName defines model for ChangeLaunchName.
type ChangeLaunchName struct {
	Name string `json:"name"`
}

// ChangeOrderDescription defines model for ChangeOrderDescription.
type ChangeOrderDescription struct {
	Description string `json:"description"`
}

// ChangeOrderName defines model for ChangeOrderName.
type ChangeOrderName struct {
	Name string `json:"name"`
}

// CreateKindRequest defines model for CreateKindRequest.
type CreateKindRequest struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
	Status      string  `json:"status"`
}

// CreateLaunchRequest defines model for CreateLaunchRequest.
type CreateLaunchRequest struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
	Status      string  `json:"status"`
}

// CreateOrderRequest defines model for CreateOrderRequest.
type CreateOrderRequest struct {
	Description *string `json:"description,omitempty"`
	Name        string  `json:"name"`
	Status      string  `json:"status"`
}

// ListKindsParams defines parameters for ListKinds.
type ListKindsParams struct {
	Cursor *string                `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit  *int                   `form:"limit,omitempty" json:"limit,omitempty"`
	Status *ListKindsParamsStatus `form:"status,omitempty" json:"status,omitempty"`
}

// ListKindsParamsStatus defines parameters for ListKinds.
type ListKindsParamsStatus string

// ListLaunchesParams defines parameters for ListLaunches.
type ListLaunchesParams struct {
	Cursor *string                   `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit  *int                      `form:"limit,omitempty" json:"limit,omitempty"`
	Status *ListLaunchesParamsStatus `form:"status,omitempty" json:"status,omitempty"`
}

// ListLaunchesParamsStatus defines parameters for ListLaunches.
type ListLaunchesParamsStatus string

// ListOrdersParams defines parameters for ListOrders.
type ListOrdersParams struct {
	Cursor *string                 `form:"cursor,omitempty" json:"cursor,omitempty"`
	Limit  *int                    `form:"limit,omitempty" json:"limit,omitempty"`
	Status *ListOrdersParamsStatus `form:"status,omitempty" json:"status,omitempty"`
}

// ListOrdersParamsStatus defines parameters for ListOrders.
type ListOrdersParamsStatus string

// CreateKindJSONRequestBody defines body for CreateKind for application/json ContentType.
type CreateKindJSONRequestBody CreateKindRequest

// ChangeKindDescriptionJSONRequestBody defines body for ChangeKindDescription for application/json ContentType.
type ChangeKindDescriptionJSONRequestBody ChangeKindDescription

// ChangeKindNameJSONRequestBody defines body for ChangeKindName for application/json ContentType.
type ChangeKindNameJSONRequestBody ChangeKindName

// CreateLaunchJSONRequestBody defines body for CreateLaunch for application/json ContentType.
type CreateLaunchJSONRequestBody CreateLaunchRequest

// ChangeLaunchDescriptionJSONRequestBody defines body for ChangeLaunchDescription for application/json ContentType.
type ChangeLaunchDescriptionJSONRequestBody ChangeLaunchDescription

// ChangeLaunchNameJSONRequestBody defines body for ChangeLaunchName for application/json ContentType.
type ChangeLaunchNameJSONRequestBody ChangeLaunchName

// CreateOrderJSONRequestBody defines body for CreateOrder for application/json ContentType.
type CreateOrderJSONRequestBody CreateOrderRequest

// ChangeOrderDescriptionJSONRequestBody defines body for ChangeOrderDescription for application/json ContentType.
type ChangeOrderDescriptionJSONRequestBody ChangeOrderDescription

// ChangeOrderNameJSONRequestBody defines body for ChangeOrderName for application/json ContentType.
type ChangeOrderNameJSONRequestBody ChangeOrderName
