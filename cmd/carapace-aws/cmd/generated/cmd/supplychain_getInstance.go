package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_getInstanceCmd = &cobra.Command{
	Use:   "get-instance",
	Short: "Enables you to programmatically retrieve the information related to an Amazon Web Services Supply Chain instance ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_getInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_getInstanceCmd).Standalone()

		supplychain_getInstanceCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier")
		supplychain_getInstanceCmd.MarkFlagRequired("instance-id")
	})
	supplychainCmd.AddCommand(supplychain_getInstanceCmd)
}
