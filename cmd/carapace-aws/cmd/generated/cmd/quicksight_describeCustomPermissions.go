package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeCustomPermissionsCmd = &cobra.Command{
	Use:   "describe-custom-permissions",
	Short: "Describes a custom permissions profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeCustomPermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeCustomPermissionsCmd).Standalone()

		quicksight_describeCustomPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the custom permissions profile that you want described.")
		quicksight_describeCustomPermissionsCmd.Flags().String("custom-permissions-name", "", "The name of the custom permissions profile to describe.")
		quicksight_describeCustomPermissionsCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeCustomPermissionsCmd.MarkFlagRequired("custom-permissions-name")
	})
	quicksightCmd.AddCommand(quicksight_describeCustomPermissionsCmd)
}
