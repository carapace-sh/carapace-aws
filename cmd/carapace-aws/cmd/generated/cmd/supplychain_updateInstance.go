package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_updateInstanceCmd = &cobra.Command{
	Use:   "update-instance",
	Short: "Enables you to programmatically update an Amazon Web Services Supply Chain instance description by providing all the relevant information such as account ID, instance ID and so on without using the AWS console.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_updateInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_updateInstanceCmd).Standalone()

		supplychain_updateInstanceCmd.Flags().String("instance-description", "", "The AWS Supply Chain instance description.")
		supplychain_updateInstanceCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier.")
		supplychain_updateInstanceCmd.Flags().String("instance-name", "", "The AWS Supply Chain instance name.")
		supplychain_updateInstanceCmd.MarkFlagRequired("instance-id")
	})
	supplychainCmd.AddCommand(supplychain_updateInstanceCmd)
}
