package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_createCustomRoutingAcceleratorCmd = &cobra.Command{
	Use:   "create-custom-routing-accelerator",
	Short: "Create a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_createCustomRoutingAcceleratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_createCustomRoutingAcceleratorCmd).Standalone()

		globalaccelerator_createCustomRoutingAcceleratorCmd.Flags().String("enabled", "", "Indicates whether an accelerator is enabled.")
		globalaccelerator_createCustomRoutingAcceleratorCmd.Flags().String("idempotency-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency—that is, the uniqueness—of the request.")
		globalaccelerator_createCustomRoutingAcceleratorCmd.Flags().String("ip-address-type", "", "The IP address type that an accelerator supports.")
		globalaccelerator_createCustomRoutingAcceleratorCmd.Flags().String("ip-addresses", "", "Optionally, if you've added your own IP address pool to Global Accelerator (BYOIP), you can choose an IPv4 address from your own pool to use for the accelerator's static IPv4 address when you create an accelerator.")
		globalaccelerator_createCustomRoutingAcceleratorCmd.Flags().String("name", "", "The name of a custom routing accelerator.")
		globalaccelerator_createCustomRoutingAcceleratorCmd.Flags().String("tags", "", "Create tags for an accelerator.")
		globalaccelerator_createCustomRoutingAcceleratorCmd.MarkFlagRequired("idempotency-token")
		globalaccelerator_createCustomRoutingAcceleratorCmd.MarkFlagRequired("name")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_createCustomRoutingAcceleratorCmd)
}
