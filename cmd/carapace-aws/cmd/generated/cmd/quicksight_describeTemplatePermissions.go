package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeTemplatePermissionsCmd = &cobra.Command{
	Use:   "describe-template-permissions",
	Short: "Describes read and write permissions on a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeTemplatePermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeTemplatePermissionsCmd).Standalone()

		quicksight_describeTemplatePermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template that you're describing.")
		quicksight_describeTemplatePermissionsCmd.Flags().String("template-id", "", "The ID for the template.")
		quicksight_describeTemplatePermissionsCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeTemplatePermissionsCmd.MarkFlagRequired("template-id")
	})
	quicksightCmd.AddCommand(quicksight_describeTemplatePermissionsCmd)
}
