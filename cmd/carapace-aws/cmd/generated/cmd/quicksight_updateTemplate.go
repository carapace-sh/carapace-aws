package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateTemplateCmd = &cobra.Command{
	Use:   "update-template",
	Short: "Updates a template from an existing Amazon Quick Sight analysis or another template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateTemplateCmd).Standalone()

	quicksight_updateTemplateCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template that you're updating.")
	quicksight_updateTemplateCmd.Flags().String("definition", "", "The definition of a template.")
	quicksight_updateTemplateCmd.Flags().String("name", "", "The name for the template.")
	quicksight_updateTemplateCmd.Flags().String("source-entity", "", "The entity that you are using as a source when you update the template.")
	quicksight_updateTemplateCmd.Flags().String("template-id", "", "The ID for the template.")
	quicksight_updateTemplateCmd.Flags().String("validation-strategy", "", "The option to relax the validation needed to update a template with definition objects.")
	quicksight_updateTemplateCmd.Flags().String("version-description", "", "A description of the current template version that is being updated.")
	quicksight_updateTemplateCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateTemplateCmd.MarkFlagRequired("template-id")
	quicksightCmd.AddCommand(quicksight_updateTemplateCmd)
}
