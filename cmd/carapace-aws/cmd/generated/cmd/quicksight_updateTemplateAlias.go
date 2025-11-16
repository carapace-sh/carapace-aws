package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateTemplateAliasCmd = &cobra.Command{
	Use:   "update-template-alias",
	Short: "Updates the template alias of a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateTemplateAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateTemplateAliasCmd).Standalone()

		quicksight_updateTemplateAliasCmd.Flags().String("alias-name", "", "The alias of the template that you want to update.")
		quicksight_updateTemplateAliasCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template alias that you're updating.")
		quicksight_updateTemplateAliasCmd.Flags().String("template-id", "", "The ID for the template.")
		quicksight_updateTemplateAliasCmd.Flags().String("template-version-number", "", "The version number of the template.")
		quicksight_updateTemplateAliasCmd.MarkFlagRequired("alias-name")
		quicksight_updateTemplateAliasCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateTemplateAliasCmd.MarkFlagRequired("template-id")
		quicksight_updateTemplateAliasCmd.MarkFlagRequired("template-version-number")
	})
	quicksightCmd.AddCommand(quicksight_updateTemplateAliasCmd)
}
