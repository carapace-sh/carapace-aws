package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_updateCustomRoutingAcceleratorCmd = &cobra.Command{
	Use:   "update-custom-routing-accelerator",
	Short: "Update a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_updateCustomRoutingAcceleratorCmd).Standalone()

	globalaccelerator_updateCustomRoutingAcceleratorCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator to update.")
	globalaccelerator_updateCustomRoutingAcceleratorCmd.Flags().String("enabled", "", "Indicates whether an accelerator is enabled.")
	globalaccelerator_updateCustomRoutingAcceleratorCmd.Flags().String("ip-address-type", "", "The IP address type that an accelerator supports.")
	globalaccelerator_updateCustomRoutingAcceleratorCmd.Flags().String("ip-addresses", "", "The IP addresses for an accelerator.")
	globalaccelerator_updateCustomRoutingAcceleratorCmd.Flags().String("name", "", "The name of the accelerator.")
	globalaccelerator_updateCustomRoutingAcceleratorCmd.MarkFlagRequired("accelerator-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_updateCustomRoutingAcceleratorCmd)
}
