package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteDefaultQbusinessApplicationCmd = &cobra.Command{
	Use:   "delete-default-qbusiness-application",
	Short: "Deletes a linked Amazon Q Business application from an Quick Sight account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteDefaultQbusinessApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteDefaultQbusinessApplicationCmd).Standalone()

		quicksight_deleteDefaultQbusinessApplicationCmd.Flags().String("aws-account-id", "", "The ID of the Quick Sight account that you want to disconnect from a Amazon Q Business application.")
		quicksight_deleteDefaultQbusinessApplicationCmd.Flags().String("namespace", "", "The Quick Sight namespace that you want to delete a linked Amazon Q Business application from.")
		quicksight_deleteDefaultQbusinessApplicationCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteDefaultQbusinessApplicationCmd)
}
