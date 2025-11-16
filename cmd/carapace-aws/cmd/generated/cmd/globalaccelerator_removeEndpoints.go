package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_removeEndpointsCmd = &cobra.Command{
	Use:   "remove-endpoints",
	Short: "Remove endpoints from an endpoint group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_removeEndpointsCmd).Standalone()

	globalaccelerator_removeEndpointsCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group.")
	globalaccelerator_removeEndpointsCmd.Flags().String("endpoint-identifiers", "", "The identifiers of the endpoints that you want to remove.")
	globalaccelerator_removeEndpointsCmd.MarkFlagRequired("endpoint-group-arn")
	globalaccelerator_removeEndpointsCmd.MarkFlagRequired("endpoint-identifiers")
	globalacceleratorCmd.AddCommand(globalaccelerator_removeEndpointsCmd)
}
