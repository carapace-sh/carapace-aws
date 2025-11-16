package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeTemplateDefinitionCmd = &cobra.Command{
	Use:   "describe-template-definition",
	Short: "Provides a detailed description of the definition of a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeTemplateDefinitionCmd).Standalone()

	quicksight_describeTemplateDefinitionCmd.Flags().String("alias-name", "", "The alias of the template that you want to describe.")
	quicksight_describeTemplateDefinitionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template.")
	quicksight_describeTemplateDefinitionCmd.Flags().String("template-id", "", "The ID of the template that you're describing.")
	quicksight_describeTemplateDefinitionCmd.Flags().String("version-number", "", "The version number of the template.")
	quicksight_describeTemplateDefinitionCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeTemplateDefinitionCmd.MarkFlagRequired("template-id")
	quicksightCmd.AddCommand(quicksight_describeTemplateDefinitionCmd)
}
