package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createTemplateAliasCmd = &cobra.Command{
	Use:   "create-template-alias",
	Short: "Creates a template alias for a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createTemplateAliasCmd).Standalone()

	quicksight_createTemplateAliasCmd.Flags().String("alias-name", "", "The name that you want to give to the template alias that you're creating.")
	quicksight_createTemplateAliasCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template that you creating an alias for.")
	quicksight_createTemplateAliasCmd.Flags().String("template-id", "", "An ID for the template.")
	quicksight_createTemplateAliasCmd.Flags().String("template-version-number", "", "The version number of the template.")
	quicksight_createTemplateAliasCmd.MarkFlagRequired("alias-name")
	quicksight_createTemplateAliasCmd.MarkFlagRequired("aws-account-id")
	quicksight_createTemplateAliasCmd.MarkFlagRequired("template-id")
	quicksight_createTemplateAliasCmd.MarkFlagRequired("template-version-number")
	quicksightCmd.AddCommand(quicksight_createTemplateAliasCmd)
}
