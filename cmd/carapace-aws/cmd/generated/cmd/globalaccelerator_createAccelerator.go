package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_createAcceleratorCmd = &cobra.Command{
	Use:   "create-accelerator",
	Short: "Create an accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_createAcceleratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_createAcceleratorCmd).Standalone()

		globalaccelerator_createAcceleratorCmd.Flags().String("enabled", "", "Indicates whether an accelerator is enabled.")
		globalaccelerator_createAcceleratorCmd.Flags().String("idempotency-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency—that is, the uniqueness—of an accelerator.")
		globalaccelerator_createAcceleratorCmd.Flags().String("ip-address-type", "", "The IP address type that an accelerator supports.")
		globalaccelerator_createAcceleratorCmd.Flags().String("ip-addresses", "", "Optionally, if you've added your own IP address pool to Global Accelerator (BYOIP), you can choose an IPv4 address from your own pool to use for the accelerator's static IPv4 address when you create an accelerator.")
		globalaccelerator_createAcceleratorCmd.Flags().String("name", "", "The name of the accelerator.")
		globalaccelerator_createAcceleratorCmd.Flags().String("tags", "", "Create tags for an accelerator.")
		globalaccelerator_createAcceleratorCmd.MarkFlagRequired("idempotency-token")
		globalaccelerator_createAcceleratorCmd.MarkFlagRequired("name")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_createAcceleratorCmd)
}
