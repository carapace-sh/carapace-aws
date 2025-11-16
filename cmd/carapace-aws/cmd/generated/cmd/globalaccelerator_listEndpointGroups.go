package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listEndpointGroupsCmd = &cobra.Command{
	Use:   "list-endpoint-groups",
	Short: "List the endpoint groups that are associated with a listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listEndpointGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_listEndpointGroupsCmd).Standalone()

		globalaccelerator_listEndpointGroupsCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
		globalaccelerator_listEndpointGroupsCmd.Flags().String("max-results", "", "The number of endpoint group objects that you want to return with this call.")
		globalaccelerator_listEndpointGroupsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		globalaccelerator_listEndpointGroupsCmd.MarkFlagRequired("listener-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_listEndpointGroupsCmd)
}
