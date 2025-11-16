package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_addEndpointsCmd = &cobra.Command{
	Use:   "add-endpoints",
	Short: "Add endpoints to an endpoint group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_addEndpointsCmd).Standalone()

	globalaccelerator_addEndpointsCmd.Flags().String("endpoint-configurations", "", "The list of endpoint objects.")
	globalaccelerator_addEndpointsCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group.")
	globalaccelerator_addEndpointsCmd.MarkFlagRequired("endpoint-configurations")
	globalaccelerator_addEndpointsCmd.MarkFlagRequired("endpoint-group-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_addEndpointsCmd)
}
