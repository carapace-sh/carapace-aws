package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateTemplatePermissionsCmd = &cobra.Command{
	Use:   "update-template-permissions",
	Short: "Updates the resource permissions for a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateTemplatePermissionsCmd).Standalone()

	quicksight_updateTemplatePermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the template.")
	quicksight_updateTemplatePermissionsCmd.Flags().String("grant-permissions", "", "A list of resource permissions to be granted on the template.")
	quicksight_updateTemplatePermissionsCmd.Flags().String("revoke-permissions", "", "A list of resource permissions to be revoked from the template.")
	quicksight_updateTemplatePermissionsCmd.Flags().String("template-id", "", "The ID for the template.")
	quicksight_updateTemplatePermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateTemplatePermissionsCmd.MarkFlagRequired("template-id")
	quicksightCmd.AddCommand(quicksight_updateTemplatePermissionsCmd)
}
