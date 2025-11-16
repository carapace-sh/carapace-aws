package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDefaultQbusinessApplicationCmd = &cobra.Command{
	Use:   "update-default-qbusiness-application",
	Short: "Updates a Amazon Q Business application that is linked to a Quick Sight account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDefaultQbusinessApplicationCmd).Standalone()

	quicksight_updateDefaultQbusinessApplicationCmd.Flags().String("application-id", "", "The ID of the Amazon Q Business application that you want to update.")
	quicksight_updateDefaultQbusinessApplicationCmd.Flags().String("aws-account-id", "", "The ID of the Quick Sight account that is connected to the Amazon Q Business application that you want to update.")
	quicksight_updateDefaultQbusinessApplicationCmd.Flags().String("namespace", "", "The Quick Sight namespace that contains the linked Amazon Q Business application.")
	quicksight_updateDefaultQbusinessApplicationCmd.MarkFlagRequired("application-id")
	quicksight_updateDefaultQbusinessApplicationCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_updateDefaultQbusinessApplicationCmd)
}
