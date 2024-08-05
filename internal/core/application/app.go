package application

import (
	"context"
	l "pub-hf-voucher-p5/external/logger"
	ps "pub-hf-voucher-p5/external/strings"
	"pub-hf-voucher-p5/internal/core/domain/broker"
	"pub-hf-voucher-p5/internal/core/domain/entity/dto"
)

type Application interface {
	GetVoucherByID(msgID string, uuid string) error
	SaveVoucher(msgID string, voucher dto.RequestVoucher) error
	UpdateVoucherByID(msgID string, id string, voucher dto.RequestVoucher) error
}

type application struct {
	ctx           context.Context
	voucherBroker broker.VoucherBroker
}

func NewApplication(ctx context.Context, voucherBroker broker.VoucherBroker) Application {
	return application{
		ctx:           ctx,
		voucherBroker: voucherBroker,
	}
}

func (app application) GetVoucherByID(msgID string, uuid string) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "GetVoucherByIDApp: ", " | ", uuid)

	inputBroker := dto.VoucherBroker{
		UUID:      uuid,
		MessageID: msgID,
	}

	if err := app.voucherBroker.GetVoucherByID(inputBroker); err != nil {
		l.Errorf(msgID, "GetVoucherByIDApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "GetVoucherByIDApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) SaveVoucher(msgID string, voucher dto.RequestVoucher) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "SaveVoucherApp: ", " | ", ps.MarshalString(voucher))

	inputBroker := dto.VoucherBroker{
		UUID:       voucher.UUID,
		MessageID:  msgID,
		Code:       voucher.Code,
		Percentage: voucher.Percentage,
		CreatedAt:  voucher.CreatedAt,
		ExpiresAt:  voucher.ExpiresAt,
	}

	if err := app.voucherBroker.SaveVoucher(inputBroker); err != nil {
		l.Errorf(msgID, "SaveVoucherApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "SaveVoucherApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) UpdateVoucherByID(msgID string, id string, voucher dto.RequestVoucher) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "UpdateVoucherByIDApp: ", " | ", id, " | ", ps.MarshalString(voucher))

	inputBroker := dto.VoucherBroker{
		UUID:       id,
		MessageID:  msgID,
		Code:       voucher.Code,
		Percentage: voucher.Percentage,
		CreatedAt:  voucher.CreatedAt,
		ExpiresAt:  voucher.ExpiresAt,
	}

	if err := app.voucherBroker.UpdateVoucherByID(inputBroker); err != nil {
		l.Errorf(msgID, "UpdateVoucherByIDApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "UpdateVoucherByIDApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) setMessageIDCtx(msgID string) {
	if app.ctx == nil {
		app.ctx = context.WithValue(context.Background(), l.MessageIDKey, msgID)
		return
	}
	app.ctx = context.WithValue(app.ctx, l.MessageIDKey, msgID)
}
