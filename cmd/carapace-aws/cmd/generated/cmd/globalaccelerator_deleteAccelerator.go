package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_deleteAcceleratorCmd = &cobra.Command{
	Use:   "delete-accelerator",
	Short: "Delete an accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_deleteAcceleratorCmd).Standalone()

	globalaccelerator_deleteAcceleratorCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of an accelerator.")
	globalaccelerator_deleteAcceleratorCmd.MarkFlagRequired("accelerator-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_deleteAcceleratorCmd)
}
