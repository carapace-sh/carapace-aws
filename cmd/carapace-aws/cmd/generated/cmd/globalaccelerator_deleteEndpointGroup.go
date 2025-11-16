package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_deleteEndpointGroupCmd = &cobra.Command{
	Use:   "delete-endpoint-group",
	Short: "Delete an endpoint group from a listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_deleteEndpointGroupCmd).Standalone()

	globalaccelerator_deleteEndpointGroupCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group to delete.")
	globalaccelerator_deleteEndpointGroupCmd.MarkFlagRequired("endpoint-group-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_deleteEndpointGroupCmd)
}
