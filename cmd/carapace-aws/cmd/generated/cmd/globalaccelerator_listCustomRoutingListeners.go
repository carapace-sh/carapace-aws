package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listCustomRoutingListenersCmd = &cobra.Command{
	Use:   "list-custom-routing-listeners",
	Short: "List the listeners for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listCustomRoutingListenersCmd).Standalone()

	globalaccelerator_listCustomRoutingListenersCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator to list listeners for.")
	globalaccelerator_listCustomRoutingListenersCmd.Flags().String("max-results", "", "The number of listener objects that you want to return with this call.")
	globalaccelerator_listCustomRoutingListenersCmd.Flags().String("next-token", "", "The token for the next set of results.")
	globalaccelerator_listCustomRoutingListenersCmd.MarkFlagRequired("accelerator-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_listCustomRoutingListenersCmd)
}
