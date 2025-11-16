package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listCustomRoutingAcceleratorsCmd = &cobra.Command{
	Use:   "list-custom-routing-accelerators",
	Short: "List the custom routing accelerators for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listCustomRoutingAcceleratorsCmd).Standalone()

	globalaccelerator_listCustomRoutingAcceleratorsCmd.Flags().String("max-results", "", "The number of custom routing Global Accelerator objects that you want to return with this call.")
	globalaccelerator_listCustomRoutingAcceleratorsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	globalacceleratorCmd.AddCommand(globalaccelerator_listCustomRoutingAcceleratorsCmd)
}
