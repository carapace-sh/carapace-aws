package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_deleteCustomRoutingAcceleratorCmd = &cobra.Command{
	Use:   "delete-custom-routing-accelerator",
	Short: "Delete a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_deleteCustomRoutingAcceleratorCmd).Standalone()

	globalaccelerator_deleteCustomRoutingAcceleratorCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the custom routing accelerator to delete.")
	globalaccelerator_deleteCustomRoutingAcceleratorCmd.MarkFlagRequired("accelerator-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_deleteCustomRoutingAcceleratorCmd)
}
