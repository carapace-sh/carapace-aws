package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateThemePermissionsCmd = &cobra.Command{
	Use:   "update-theme-permissions",
	Short: "Updates the resource permissions for a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateThemePermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateThemePermissionsCmd).Standalone()

		quicksight_updateThemePermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme.")
		quicksight_updateThemePermissionsCmd.Flags().String("grant-permissions", "", "A list of resource permissions to be granted for the theme.")
		quicksight_updateThemePermissionsCmd.Flags().String("revoke-permissions", "", "A list of resource permissions to be revoked from the theme.")
		quicksight_updateThemePermissionsCmd.Flags().String("theme-id", "", "The ID for the theme.")
		quicksight_updateThemePermissionsCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateThemePermissionsCmd.MarkFlagRequired("theme-id")
	})
	quicksightCmd.AddCommand(quicksight_updateThemePermissionsCmd)
}
