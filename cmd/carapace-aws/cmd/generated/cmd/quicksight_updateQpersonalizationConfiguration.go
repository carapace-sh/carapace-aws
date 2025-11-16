package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateQpersonalizationConfigurationCmd = &cobra.Command{
	Use:   "update-qpersonalization-configuration",
	Short: "Updates a personalization configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateQpersonalizationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateQpersonalizationConfigurationCmd).Standalone()

		quicksight_updateQpersonalizationConfigurationCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account account that contains the personalization configuration that the user wants to update.")
		quicksight_updateQpersonalizationConfigurationCmd.Flags().String("personalization-mode", "", "An option to allow Amazon Quick Sight to customize data stories with user specific metadata, specifically location and job information, in your IAM Identity Center instance.")
		quicksight_updateQpersonalizationConfigurationCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateQpersonalizationConfigurationCmd.MarkFlagRequired("personalization-mode")
	})
	quicksightCmd.AddCommand(quicksight_updateQpersonalizationConfigurationCmd)
}
