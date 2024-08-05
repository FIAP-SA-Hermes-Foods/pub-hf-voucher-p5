package broker

import (
	"pub-hf-voucher-p5/internal/core/domain/entity/dto"
)

type VoucherBroker interface {
	GetVoucherByID(input dto.VoucherBroker) error
	SaveVoucher(input dto.VoucherBroker) error
	UpdateVoucherByID(input dto.VoucherBroker) error
}
