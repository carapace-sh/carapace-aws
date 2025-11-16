package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeThemePermissionsCmd = &cobra.Command{
	Use:   "describe-theme-permissions",
	Short: "Describes the read and write permissions for a theme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeThemePermissionsCmd).Standalone()

	quicksight_describeThemePermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme that you're describing.")
	quicksight_describeThemePermissionsCmd.Flags().String("theme-id", "", "The ID for the theme that you want to describe permissions for.")
	quicksight_describeThemePermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeThemePermissionsCmd.MarkFlagRequired("theme-id")
	quicksightCmd.AddCommand(quicksight_describeThemePermissionsCmd)
}
