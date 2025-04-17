package handlers

import (
	pb "cloud_commons/order"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"orders/models"
	"orders/services"
)

type OrderHandler struct {
	service *services.OrderService
	pb.UnimplementedOrderServiceServer
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{service: services.NewOrderService()}
}

func (h *OrderHandler) Save(ctx context.Context, req *pb.Order) (*pb.Order, error) {

	order := models.Order{
		Id:        req.Id,
		Status:    req.Status,
		Quantity:  req.Quantity,
		ProductID: req.ProductId,
	}

	savedOrder, err := h.service.SaveOrder(&order)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Failed to save the order: %v",
			err.Error())
	}

	return &pb.Order{
		Id:        savedOrder.Id,
		Status:    savedOrder.Status,
		Quantity:  savedOrder.Quantity,
		ProductId: savedOrder.ProductID,
	}, nil
}

func (h *OrderHandler) Update(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "Order ID is required")
	}

	order := models.Order{
		Id:        req.Id,
		Status:    req.Status,
		Quantity:  req.Quantity,
		ProductID: req.ProductId,
	}

	updatedOrder, err := h.service.UpdateOrder(req.Id, &order)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Failed to update the order: %v",
			err.Error())
	}

	return &pb.Order{
		Id:        updatedOrder.Id,
		Status:    updatedOrder.Status,
		Quantity:  updatedOrder.Quantity,
		ProductId: updatedOrder.ProductID,
	}, nil
}

func (h *OrderHandler) Get(ctx context.Context, req *pb.OrderId) (*pb.Order, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "Order ID is required")
	}

	order, err := h.service.GetOrder(req.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Failed to get the order: %v",
			err.Error())
	}

	return &pb.Order{
		Id:        order.Id,
		Status:    order.Status,
		Quantity:  order.Quantity,
		ProductId: order.ProductID,
	}, nil
}

func (h *OrderHandler) List(req *pb.NoParams, stream pb.OrderService_ListServer) error {

	orders, err := h.service.ListOrder()
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Failed to list the orders: %v",
			err.Error())
	}

	for _, order := range orders {
		pbOrder := &pb.Order{
			Id:        order.Id,
			Status:    order.Status,
			Quantity:  order.Quantity,
			ProductId: order.ProductID,
		}

		if err := stream.Send(pbOrder); err != nil {
			return status.Errorf(
				codes.Internal,
				"Failed to send order: %v",
				err.Error())
		}
	}

	return nil
}

func (h *OrderHandler) Delete(ctx context.Context, req *pb.OrderId) (*pb.DeleteResponse, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "Order ID is required")
	}

	err := h.service.DeleteOrder(req.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Failed to delete the order: %v",
			err.Error())
	}

	return &pb.DeleteResponse{
		Message: "Successfully deleted the order",
	}, nil
}
