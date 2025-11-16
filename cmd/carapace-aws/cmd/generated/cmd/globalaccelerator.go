package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalacceleratorCmd = &cobra.Command{
	Use:   "globalaccelerator",
	Short: "Global Accelerator\n\nThis is the *Global Accelerator API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalacceleratorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalacceleratorCmd).Standalone()

	})
	rootCmd.AddCommand(globalacceleratorCmd)
}
