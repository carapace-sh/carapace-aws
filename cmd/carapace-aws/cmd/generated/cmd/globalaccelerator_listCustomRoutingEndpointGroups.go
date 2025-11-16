package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listCustomRoutingEndpointGroupsCmd = &cobra.Command{
	Use:   "list-custom-routing-endpoint-groups",
	Short: "List the endpoint groups that are associated with a listener for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listCustomRoutingEndpointGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_listCustomRoutingEndpointGroupsCmd).Standalone()

		globalaccelerator_listCustomRoutingEndpointGroupsCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener to list endpoint groups for.")
		globalaccelerator_listCustomRoutingEndpointGroupsCmd.Flags().String("max-results", "", "The number of endpoint group objects that you want to return with this call.")
		globalaccelerator_listCustomRoutingEndpointGroupsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		globalaccelerator_listCustomRoutingEndpointGroupsCmd.MarkFlagRequired("listener-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_listCustomRoutingEndpointGroupsCmd)
}
