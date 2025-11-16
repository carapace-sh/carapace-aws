package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeTemplateCmd = &cobra.Command{
	Use:   "describe-template",
	Short: "Describes a template's metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeTemplateCmd).Standalone()

	quicksight_describeTemplateCmd.Flags().String("alias-name", "", "The alias of the template that you want to describe.")
	quicksight_describeTemplateCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template that you're describing.")
	quicksight_describeTemplateCmd.Flags().String("template-id", "", "The ID for the template.")
	quicksight_describeTemplateCmd.Flags().String("version-number", "", "(Optional) The number for the version to describe.")
	quicksight_describeTemplateCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeTemplateCmd.MarkFlagRequired("template-id")
	quicksightCmd.AddCommand(quicksight_describeTemplateCmd)
}
