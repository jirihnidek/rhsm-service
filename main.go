package main

import (
	"context"
	"fmt"
	"github.com/jirihnidek/rhsm-service/interface/com_redhat_rhsm"
	"github.com/jirihnidek/rhsm-service/interface/com_redhat_rhsm/consumer"
	"github.com/jirihnidek/rhsm2"
	"github.com/rs/zerolog/log"
	"github.com/varlink/go/varlink"
)

type ComRedHatRHSM struct {
	comredhatrhsm.VarlinkInterface
}

func (*ComRedHatRHSM) Ping(ctx context.Context, call comredhatrhsm.VarlinkCall, locale string) error {
	log.Debug().Msgf("method Ping(%s) called", locale)
	rhsmClient, err := rhsm2.GetRHSMClient(nil)
	if err != nil {
		return fmt.Errorf("unable to create rhsm client: %s", err)
	}

	_, err = rhsmClient.GetServerEndpoints(nil)
	if err != nil {
		log.Debug().Msgf("server not running")
		return call.ReplyPing(ctx, false)
	}
	log.Debug().Msgf("server running")
	return call.ReplyPing(ctx, true)
}

func main() {
	rootInterface := ComRedHatRHSM{}
	consumerInterface := ComRedHatRHSMConsumer{}

	service, _ := varlink.NewService(
		"Red Hat",
		"RHSM",
		"1",
		"https://github.com/jirirhidek/rhsm-service/",
	)

	ctx := context.TODO()

	err := service.RegisterInterface(comredhatrhsm.VarlinkNew(&rootInterface))
	if err != nil {
		panic(err)
	}

	err = service.RegisterInterface(comredhatrhsmconsumer.VarlinkNew(&consumerInterface))
	if err != nil {
		panic(err)
	}

	sockAddress := "unix:/run/com.redhat.rhsm"
	log.Debug().Msgf("listening on %s...", sockAddress)
	err = service.Listen(ctx, sockAddress, 0)

	if err != nil {
		panic(err)
	}
}
