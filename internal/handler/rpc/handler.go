package rpc

import (
	"context"
	l "pub-hf-voucher-p5/external/logger"
	"pub-hf-voucher-p5/internal/core/application"
	"pub-hf-voucher-p5/internal/core/domain/entity/dto"
	cp "pub-hf-voucher-p5/voucher_pub_proto"
	"strconv"
)

type HandlerGRPC interface {
	Handler() *handlerGRPC
}

type handlerGRPC struct {
	app application.Application
	cp.UnimplementedVoucherServer
}

func NewHandler(app application.Application) HandlerGRPC {
	return &handlerGRPC{app: app}
}

func (h *handlerGRPC) Handler() *handlerGRPC {
	return h
}

func (h *handlerGRPC) GetVoucherByID(ctx context.Context, req *cp.GetVoucherByIDRequest) (*cp.GetVoucherByIDResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	if err := h.app.GetVoucherByID(msgID, req.Uuid); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handlerGRPC) CreateVoucher(ctx context.Context, req *cp.CreateVoucherRequest) (*cp.CreateVoucherResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	input := dto.RequestVoucher{
		Code:       req.Code,
		Percentage: strconv.FormatInt(req.Percentage, 10),
		ExpiresAt:  req.ExpiresAt,
	}

	if err := h.app.SaveVoucher(msgID, input); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *handlerGRPC) UpdateVoucher(ctx context.Context, req *cp.UpdateVoucherByIDRequest) (*cp.UpdateVoucherByIDResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	input := dto.RequestVoucher{
		Code:       req.Code,
		Percentage: strconv.FormatInt(req.Percentage, 10),
		CreatedAt:  req.CreatedAt,
		ExpiresAt:  req.ExpiresAt,
	}

	if err := h.app.UpdateVoucherByID(msgID, req.Uuid, input); err != nil {
		return nil, err
	}

	return nil, nil
}
