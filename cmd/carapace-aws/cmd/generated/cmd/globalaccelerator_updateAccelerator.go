package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_updateAcceleratorCmd = &cobra.Command{
	Use:   "update-accelerator",
	Short: "Update an accelerator to make changes, such as the following:\n\n- Change the name of the accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_updateAcceleratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_updateAcceleratorCmd).Standalone()

		globalaccelerator_updateAcceleratorCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator to update.")
		globalaccelerator_updateAcceleratorCmd.Flags().String("enabled", "", "Indicates whether an accelerator is enabled.")
		globalaccelerator_updateAcceleratorCmd.Flags().String("ip-address-type", "", "The IP address type that an accelerator supports.")
		globalaccelerator_updateAcceleratorCmd.Flags().String("ip-addresses", "", "The IP addresses for an accelerator.")
		globalaccelerator_updateAcceleratorCmd.Flags().String("name", "", "The name of the accelerator.")
		globalaccelerator_updateAcceleratorCmd.MarkFlagRequired("accelerator-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_updateAcceleratorCmd)
}
