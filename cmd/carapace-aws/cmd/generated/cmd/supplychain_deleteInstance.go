package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_deleteInstanceCmd = &cobra.Command{
	Use:   "delete-instance",
	Short: "Enables you to programmatically delete an Amazon Web Services Supply Chain instance by deleting the KMS keys and relevant information associated with the API without using the Amazon Web Services console.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_deleteInstanceCmd).Standalone()

	supplychain_deleteInstanceCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier.")
	supplychain_deleteInstanceCmd.MarkFlagRequired("instance-id")
	supplychainCmd.AddCommand(supplychain_deleteInstanceCmd)
}
