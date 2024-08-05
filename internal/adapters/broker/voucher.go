package broker

import (
	l "pub-hf-voucher-p5/external/logger"
	ps "pub-hf-voucher-p5/external/strings"
	sqsBroker "pub-hf-voucher-p5/internal/core/broker"
	vBroker "pub-hf-voucher-p5/internal/core/domain/broker"
	"pub-hf-voucher-p5/internal/core/domain/entity/dto"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var _ vBroker.VoucherBroker = (*voucherBroker)(nil)

type voucherBroker struct {
	queueURL string
	broker   sqsBroker.SQSBroker
}

func NewVoucherBroker(broker sqsBroker.SQSBroker, queueURL string) *voucherBroker {
	return &voucherBroker{broker: broker, queueURL: queueURL}
}

func (v *voucherBroker) GetVoucherByID(input dto.VoucherBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &v.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := v.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (v *voucherBroker) SaveVoucher(input dto.VoucherBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &v.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := v.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (v *voucherBroker) UpdateVoucherByID(input dto.VoucherBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &v.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := v.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

