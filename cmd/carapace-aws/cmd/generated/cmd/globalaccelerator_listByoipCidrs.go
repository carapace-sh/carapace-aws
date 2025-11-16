package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listByoipCidrsCmd = &cobra.Command{
	Use:   "list-byoip-cidrs",
	Short: "Lists the IP address ranges that were specified in calls to [ProvisionByoipCidr](https://docs.aws.amazon.com/global-accelerator/latest/api/ProvisionByoipCidr.html), including the current state and a history of state changes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listByoipCidrsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_listByoipCidrsCmd).Standalone()

		globalaccelerator_listByoipCidrsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		globalaccelerator_listByoipCidrsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_listByoipCidrsCmd)
}
