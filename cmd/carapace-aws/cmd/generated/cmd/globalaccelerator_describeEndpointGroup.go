package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_describeEndpointGroupCmd = &cobra.Command{
	Use:   "describe-endpoint-group",
	Short: "Describe an endpoint group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_describeEndpointGroupCmd).Standalone()

	globalaccelerator_describeEndpointGroupCmd.Flags().String("endpoint-group-arn", "", "The Amazon Resource Name (ARN) of the endpoint group to describe.")
	globalaccelerator_describeEndpointGroupCmd.MarkFlagRequired("endpoint-group-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_describeEndpointGroupCmd)
}
