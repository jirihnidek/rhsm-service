package main

import (
	"context"
	"fmt"
	comredhatrhsmconsumer "github.com/jirihnidek/rhsm-service/interface/com_redhat_rhsm/consumer"
	"github.com/jirihnidek/rhsm2"
	"github.com/rs/zerolog/log"
)

type ComRedHatRHSMConsumer struct {
	comredhatrhsmconsumer.VarlinkInterface
	consumerCertFileName string
	consumerKeyFileName  string
	uuid                 string
	orgID                string
}

func (consumer *ComRedHatRHSMConsumer) GetUUID(
	ctx context.Context,
	call comredhatrhsmconsumer.VarlinkCall,
	locale string,
) error {
	log.Debug().Msgf("method GetUUID(%s) called", locale)
	err := consumer.readConsumerData()
	if err != nil {
		return err
	}
	return call.ReplyGetUUID(ctx, consumer.uuid)
}

func (consumer *ComRedHatRHSMConsumer) GetOrg(
	ctx context.Context,
	call comredhatrhsmconsumer.VarlinkCall,
	locale string,
) error {
	log.Debug().Msgf("method GetOrg(%s) called", locale)
	err := consumer.readConsumerData()
	if err != nil {
		return err
	}
	return call.ReplyGetUUID(ctx, consumer.orgID)
}

// readConsumerData tries to read data from consumer certificate
// and set fields in rhsm2 structure
func (consumer *ComRedHatRHSMConsumer) readConsumerData() error {
	rhsmClient, err := rhsm2.GetRHSMClient(nil)
	if err != nil {
		return fmt.Errorf("unable to create rhsm client: %s", err)
	}

	log.Debug().Msgf("reloading consumer certificate...")
	consumerUUID, err := rhsmClient.GetConsumerUUID()
	if err != nil {
		consumer.uuid = ""
	} else {
		consumer.uuid = *consumerUUID
	}

	ownerId, err := rhsmClient.GetOwner()
	if err != nil {
		consumer.orgID = ""
	} else {
		consumer.orgID = *ownerId
	}

	return nil
}
