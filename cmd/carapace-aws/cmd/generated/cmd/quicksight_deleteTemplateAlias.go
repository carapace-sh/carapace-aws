package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteTemplateAliasCmd = &cobra.Command{
	Use:   "delete-template-alias",
	Short: "Deletes the item that the specified template alias points to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteTemplateAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteTemplateAliasCmd).Standalone()

		quicksight_deleteTemplateAliasCmd.Flags().String("alias-name", "", "The name for the template alias.")
		quicksight_deleteTemplateAliasCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the item to delete.")
		quicksight_deleteTemplateAliasCmd.Flags().String("template-id", "", "The ID for the template that the specified alias is for.")
		quicksight_deleteTemplateAliasCmd.MarkFlagRequired("alias-name")
		quicksight_deleteTemplateAliasCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteTemplateAliasCmd.MarkFlagRequired("template-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteTemplateAliasCmd)
}
