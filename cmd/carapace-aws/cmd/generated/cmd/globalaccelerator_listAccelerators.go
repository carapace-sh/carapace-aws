package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listAcceleratorsCmd = &cobra.Command{
	Use:   "list-accelerators",
	Short: "List the accelerators for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listAcceleratorsCmd).Standalone()

	globalaccelerator_listAcceleratorsCmd.Flags().String("max-results", "", "The number of Global Accelerator objects that you want to return with this call.")
	globalaccelerator_listAcceleratorsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	globalacceleratorCmd.AddCommand(globalaccelerator_listAcceleratorsCmd)
}
