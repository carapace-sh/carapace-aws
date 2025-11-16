package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteThemeAliasCmd = &cobra.Command{
	Use:   "delete-theme-alias",
	Short: "Deletes the version of the theme that the specified theme alias points to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteThemeAliasCmd).Standalone()

	quicksight_deleteThemeAliasCmd.Flags().String("alias-name", "", "The unique name for the theme alias to delete.")
	quicksight_deleteThemeAliasCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the theme alias to delete.")
	quicksight_deleteThemeAliasCmd.Flags().String("theme-id", "", "The ID for the theme that the specified alias is for.")
	quicksight_deleteThemeAliasCmd.MarkFlagRequired("alias-name")
	quicksight_deleteThemeAliasCmd.MarkFlagRequired("aws-account-id")
	quicksight_deleteThemeAliasCmd.MarkFlagRequired("theme-id")
	quicksightCmd.AddCommand(quicksight_deleteThemeAliasCmd)
}
