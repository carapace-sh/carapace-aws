package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeAccountCustomPermissionCmd = &cobra.Command{
	Use:   "describe-account-custom-permission",
	Short: "Describes the custom permissions profile that is applied to an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeAccountCustomPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeAccountCustomPermissionCmd).Standalone()

		quicksight_describeAccountCustomPermissionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account for which you want to describe the applied custom permissions profile.")
		quicksight_describeAccountCustomPermissionCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_describeAccountCustomPermissionCmd)
}
