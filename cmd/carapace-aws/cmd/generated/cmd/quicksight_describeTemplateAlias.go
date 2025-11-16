package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeTemplateAliasCmd = &cobra.Command{
	Use:   "describe-template-alias",
	Short: "Describes the template alias for a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeTemplateAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeTemplateAliasCmd).Standalone()

		quicksight_describeTemplateAliasCmd.Flags().String("alias-name", "", "The name of the template alias that you want to describe.")
		quicksight_describeTemplateAliasCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template alias that you're describing.")
		quicksight_describeTemplateAliasCmd.Flags().String("template-id", "", "The ID for the template.")
		quicksight_describeTemplateAliasCmd.MarkFlagRequired("alias-name")
		quicksight_describeTemplateAliasCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeTemplateAliasCmd.MarkFlagRequired("template-id")
	})
	quicksightCmd.AddCommand(quicksight_describeTemplateAliasCmd)
}
