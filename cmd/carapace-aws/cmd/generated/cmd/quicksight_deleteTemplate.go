package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteTemplateCmd = &cobra.Command{
	Use:   "delete-template",
	Short: "Deletes a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteTemplateCmd).Standalone()

	quicksight_deleteTemplateCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template that you're deleting.")
	quicksight_deleteTemplateCmd.Flags().String("template-id", "", "An ID for the template you want to delete.")
	quicksight_deleteTemplateCmd.Flags().String("version-number", "", "Specifies the version of the template that you want to delete.")
	quicksight_deleteTemplateCmd.MarkFlagRequired("aws-account-id")
	quicksight_deleteTemplateCmd.MarkFlagRequired("template-id")
	quicksightCmd.AddCommand(quicksight_deleteTemplateCmd)
}
