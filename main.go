package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Model struct {
	provider           string // the provider's pubkey
	service            string // ie btc-mainnet-fullnode TODO: make this an enum and select from list
	cType              int    // contract type, 0 is subscription, 1 is pay-as-you-go
	deposit            uint64 // deposit amount. Subscriptions should make sense in that duration and rate equal deposit
	duration           uint64 // number of blocks to make a subscription. There are lower and higher limits to this number
	rate               uint64 // should equal the provider's rate
	settlementDuration uint64 // this number should equal the same number
	qpm                uint64 // this set the rate limit for the contract. A higher value will come with a higher cost
	auth               int    // defines if the contract has strict authorization (0)  or open (1)
	delegate           string // if you'd like to have a different private key spend from the contract, put that pubkey here
	err                error
}

var model Model

func main() {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the provider's pubkey").
				Value(&model.provider),

			huh.NewSelect[string]().
				Title("Select a service").
				Options(
					huh.NewOption("Bitcoin Mainnet", "btc-mainnet-fullnode"),
				).
				Value(&model.service),

			huh.NewSelect[int]().
				Title("Select a contract type").
				Options(
					huh.NewOption("Subscription", 0),
					huh.NewOption("Pay-as-you-go", 1),
				).
				Value(&model.cType),

			huh.NewInput().
				Title("Enter the deposit amount").
				Validate(func(input string) error {
					value, err := strconv.ParseUint(input, 10, 64)
					if err != nil {
						model.err = err
						return err
					}
					model.deposit = value
					return nil
				}),

			huh.NewInput().
				Title("Enter the number of blocks to make a subscription.").
				Validate(func(input string) error {
					value, err := strconv.ParseUint(input, 10, 64)
					if err != nil {
						model.err = err
						return err
					}
					model.duration = value
					return nil
				}),

			huh.NewInput().
				Title("Enter the provider's rate").
				Validate(func(input string) error {
					value, err := strconv.ParseUint(input, 10, 64)
					if err != nil {
						model.err = err
						return err
					}
					model.rate = value
					return nil
				}),

			huh.NewInput().
				Title("Enter the settlement duration").
				Validate(func(input string) error {
					value, err := strconv.ParseUint(input, 10, 64)
					if err != nil {
						model.err = err
						return err
					}
					model.settlementDuration = value
					return nil
				}),

			huh.NewInput().
				Title("Enter the rate limit for the contract").
				Validate(func(input string) error {
					value, err := strconv.ParseUint(input, 10, 64)
					if err != nil {
						model.err = err
						return err
					}
					model.qpm = value
					return nil
				}),

			huh.NewSelect[int]().
				Title("Select the authorization type").
				Options(
					huh.NewOption("Strict", 0),
					huh.NewOption("Open", 1),
				).
				Value(&model.auth),

			huh.NewInput().
				Title("Enter the delegate pubkey (optional)").
				Value(&model.delegate),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if model.err != nil {
		log.Fatal(model.err)
	}

	command := fmt.Sprintf(`arkeod tx arkeo open-contract -y --from <user> --keyring-backend file --node "tcp://seed.arkeo.network:26657" -- "%s" "%s" "<your pubkey>" "%d" "%d" "%d" "%d" "%d" "%d" "%d" "%s"`,
		model.provider,
		model.service,
		model.cType,
		model.deposit,
		model.duration,
		model.rate,
		model.qpm,
		model.settlementDuration,
		model.auth,
		model.delegate,
	)

	fmt.Println(command)

	// execute the command
}
