package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Short: "Creates a template either from a `TemplateDefinition` or from an existing Quick Sight analysis or template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_createTemplateCmd).Standalone()

		quicksight_createTemplateCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
		quicksight_createTemplateCmd.Flags().String("definition", "", "The definition of a template.")
		quicksight_createTemplateCmd.Flags().String("name", "", "A display name for the template.")
		quicksight_createTemplateCmd.Flags().String("permissions", "", "A list of resource permissions to be set on the template.")
		quicksight_createTemplateCmd.Flags().String("source-entity", "", "The entity that you are using as a source when you create the template.")
		quicksight_createTemplateCmd.Flags().String("tags", "", "Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.")
		quicksight_createTemplateCmd.Flags().String("template-id", "", "An ID for the template that you want to create.")
		quicksight_createTemplateCmd.Flags().String("validation-strategy", "", "TThe option to relax the validation needed to create a template with definition objects.")
		quicksight_createTemplateCmd.Flags().String("version-description", "", "A description of the current template version being created.")
		quicksight_createTemplateCmd.MarkFlagRequired("aws-account-id")
		quicksight_createTemplateCmd.MarkFlagRequired("template-id")
	})
	quicksightCmd.AddCommand(quicksight_createTemplateCmd)
}
