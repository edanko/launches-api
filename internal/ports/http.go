package ports

import (
	"math"
	"net/http"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/edanko/launches-api/internal/app"
	commands2 "github.com/edanko/launches-api/internal/app/commands"
	queries2 "github.com/edanko/launches-api/internal/app/queries"
	"github.com/edanko/launches-api/pkg/pagination"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(application app.Application) HTTPServer {
	return HTTPServer{
		app: application,
	}
}

func (h HTTPServer) ListLaunches(c echo.Context, params ListLaunchesParams) error {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) CreateLaunch(c echo.Context) error {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) DeleteLaunch(c echo.Context, id openapi_types.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) GetLaunch(c echo.Context, id openapi_types.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeLaunchDescription(c echo.Context, id openapi_types.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ChangeLaunchName(c echo.Context, id openapi_types.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeLaunchDraft(c echo.Context, id openapi_types.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) MakeLaunchPublished(c echo.Context, id openapi_types.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (h HTTPServer) ListOrders(c echo.Context, params ListOrdersParams) error {
	var limitInt int

	if params.Limit != nil {
		limitInt = *params.Limit
	}
	limit := int(math.Max(float64(limitInt), 10))

	var createdAt *time.Time
	var id *uuid.UUID
	if params.Cursor != nil {
		createdAtCursor, idCursor, err := pagination.DecodeCursor(*params.Cursor)
		if err != nil {
			// httperr.RespondWithSlugError(err, w, r)
			return err
		}

		createdAt = &createdAtCursor
		id = &idCursor
	}

	appOrders, err := h.app.Queries.ListOrders.Handle(
		c.Request().Context(),
		queries2.ListOrdersRequest{
			Status:    (*string)(params.Status),
			Limit:     &limit,
			CreatedAt: createdAt,
			ID:        id,
		},
	)
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	var nextCursor string
	if len(appOrders) > 0 {
		nextCursor = pagination.EncodeCursor(appOrders[len(appOrders)-1].CreatedAt, appOrders[len(appOrders)-1].ID)
	}

	ordersResp := ListOrdersResponse{
		NextCursor: nextCursor,
		Orders:     appOrdersToResponse(appOrders),
	}
	// render.Respond(w, r, ordersResp)
	return c.JSON(http.StatusOK, ordersResp)

}

func (h HTTPServer) CreateOrder(c echo.Context) error {
	postOrder := CreateOrderRequest{}
	if err := c.Bind(&postOrder); err != nil {
		// httperr.BadRequest("invalid-request", err, w, r)
		return err
	}

	cmd := commands2.CreateOrder{
		ID:          uuid.New(),
		Name:        postOrder.Name,
		Description: postOrder.Description,
		Status:      postOrder.Status,
	}

	err := h.app.CommandBus.Send(c.Request().Context(), cmd)
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Header().Set("content-location", "/orders/"+cmd.ID.String())
	c.Response().Status = http.StatusNoContent
	return nil

}

func (h HTTPServer) DeleteOrder(c echo.Context, id openapi_types.UUID) error {
	err := h.app.CommandBus.Send(c.Request().Context(), commands2.DeleteOrder{
		ID: id,
	})
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}
	return nil

}

func (h HTTPServer) GetOrder(c echo.Context, id openapi_types.UUID) error {
	appOrder, err := h.app.Queries.GetOrder.Handle(c.Request().Context(), queries2.GetOrderRequest{
		ID: id,
	})
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	orderResp := GetOrderResponse{
		Id:          appOrder.ID,
		Name:        appOrder.Name,
		Description: appOrder.Description,
		Status:      OrderStatus(appOrder.Status),
	}
	return c.JSON(http.StatusOK, orderResp)
	// render.Respond(w, r, orderResp)
}

func (h HTTPServer) ChangeOrderDescription(c echo.Context, id openapi_types.UUID) error {
	postOrder := ChangeOrderDescription{}
	if err := c.Bind(&postOrder); err != nil {
		// httperr.BadRequest("invalid-request", err, w, r)
		return err
	}

	cmd := commands2.ChangeOrderDescription{
		ID:          id,
		Description: postOrder.Description,
	}

	err := h.app.CommandBus.Send(c.Request().Context(), cmd)
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Status = http.StatusNoContent
	return nil

}

func (h HTTPServer) ChangeOrderName(c echo.Context, id openapi_types.UUID) error {
	postOrder := ChangeOrderName{}
	if err := c.Bind(&postOrder); err != nil {
		// httperr.BadRequest("invalid-request", err, w, r)
		return err
	}

	cmd := commands2.ChangeOrderName{
		ID:   id,
		Name: postOrder.Name,
	}

	err := h.app.CommandBus.Send(c.Request().Context(), cmd)
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Status = http.StatusNoContent
	return nil
}

func (h HTTPServer) MakeOrderDraft(c echo.Context, id openapi_types.UUID) error {
	err := h.app.CommandBus.Send(c.Request().Context(), commands2.MakeOrderDraft{
		ID: id,
	})
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Status = http.StatusNoContent
	return nil
}

func (h HTTPServer) MakeOrderPublished(c echo.Context, id openapi_types.UUID) error {
	err := h.app.CommandBus.Send(c.Request().Context(), commands2.MakeOrderPublished{
		ID: id,
	})
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Status = http.StatusNoContent
	return nil
}

func (h HTTPServer) GetKind(c echo.Context, id uuid.UUID) error {
	appKind, err := h.app.Queries.GetKind.Handle(c.Request().Context(), queries2.GetKindRequest{
		ID: id,
	})
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	kindResp := GetKindResponse{
		Id:          appKind.ID,
		Name:        appKind.Name,
		Description: appKind.Description,
		Status:      KindStatus(appKind.Status),
	}
	return c.JSON(http.StatusOK, kindResp)
}

func (h HTTPServer) ListKinds(c echo.Context, params ListKindsParams) error {
	var limitInt int

	if params.Limit != nil {
		limitInt = *params.Limit
	}
	limit := int(math.Max(float64(limitInt), 10))

	var createdAt *time.Time
	var id *uuid.UUID
	if params.Cursor != nil {
		createdAtCursor, idCursor, err := pagination.DecodeCursor(*params.Cursor)
		if err != nil {
			// httperr.RespondWithSlugError(err, w, r)
			return err
		}

		createdAt = &createdAtCursor
		id = &idCursor
	}

	appKinds, err := h.app.Queries.ListKinds.Handle(
		c.Request().Context(),
		queries2.ListKindsRequest{
			Status:    (*string)(params.Status),
			Limit:     &limit,
			CreatedAt: createdAt,
			ID:        id,
		},
	)
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	var nextCursor string
	if len(appKinds) > 0 {
		nextCursor = pagination.EncodeCursor(appKinds[len(appKinds)-1].CreatedAt, appKinds[len(appKinds)-1].ID)
	}

	kindsResp := ListKindsResponse{
		NextCursor: nextCursor,
		Kinds:      appKindsToResponse(appKinds),
	}
	// render.Respond(w, r, kindsResp)
	return c.JSON(http.StatusOK, kindsResp)

}

func (h HTTPServer) CreateKind(c echo.Context) error {
	postKind := CreateKindRequest{}
	if err := c.Bind(&postKind); err != nil {
		// httperr.BadRequest("invalid-request", err, w, r)
		return err
	}

	cmd := commands2.CreateKind{
		ID:          uuid.New(),
		Name:        postKind.Name,
		Description: postKind.Description,
		Status:      postKind.Status,
	}

	err := h.app.CommandBus.Send(c.Request().Context(), cmd)
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Header().Set("content-location", "/kinds/"+cmd.ID.String())
	c.Response().Status = http.StatusNoContent
	return nil
}

func (h HTTPServer) DeleteKind(c echo.Context, id uuid.UUID) error {
	err := h.app.CommandBus.Send(c.Request().Context(), commands2.DeleteKind{
		ID: id,
	})
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}
	return nil
}

func (h HTTPServer) MakeKindDraft(c echo.Context, id uuid.UUID) error {
	err := h.app.CommandBus.Send(c.Request().Context(), commands2.MakeKindDraft{
		ID: id,
	})
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Status = http.StatusNoContent
	return nil
}

func (h HTTPServer) MakeKindPublished(c echo.Context, id uuid.UUID) error {
	err := h.app.CommandBus.Send(c.Request().Context(), commands2.MakeKindPublished{
		ID: id,
	})
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Status = http.StatusNoContent
	return nil
}

func (h HTTPServer) ChangeKindDescription(c echo.Context, id uuid.UUID) error {
	postKind := ChangeKindDescription{}
	if err := c.Bind(&postKind); err != nil {
		// httperr.BadRequest("invalid-request", err, w, r)
		return err
	}

	cmd := commands2.ChangeKindDescription{
		ID:          id,
		Description: postKind.Description,
	}

	err := h.app.CommandBus.Send(c.Request().Context(), cmd)
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Status = http.StatusNoContent

	return nil
	// w.WriteHeader(http.StatusNoContent)
}

func (h HTTPServer) ChangeKindName(c echo.Context, id uuid.UUID) error {
	postKind := ChangeKindName{}
	if err := c.Bind(&postKind); err != nil {
		// httperr.BadRequest("invalid-request", err, w, r)
		return err
	}

	cmd := commands2.ChangeKindName{
		ID:   id,
		Name: postKind.Name,
	}

	err := h.app.CommandBus.Send(c.Request().Context(), cmd)
	if err != nil {
		// httperr.RespondWithSlugError(err, w, r)
		return err
	}

	c.Response().Status = http.StatusNoContent

	// w.WriteHeader(http.StatusNoContent)
	return nil

}

func appKindsToResponse(appKinds []queries2.Kind) []Kind {
	kinds := make([]Kind, 0, len(appKinds))
	for _, km := range appKinds {
		k := Kind{
			Id:          km.ID,
			Name:        km.Name,
			Description: km.Description,
			Status:      KindStatus(km.Status),
		}

		kinds = append(kinds, k)
	}

	return kinds
}

func appOrdersToResponse(appOrders []queries2.Order) []Order {
	orders := make([]Order, 0, len(appOrders))
	for _, km := range appOrders {
		k := Order{
			Id:          km.ID,
			Name:        km.Name,
			Description: km.Description,
			Status:      OrderStatus(km.Status),
		}

		orders = append(orders, k)
	}

	return orders
}
