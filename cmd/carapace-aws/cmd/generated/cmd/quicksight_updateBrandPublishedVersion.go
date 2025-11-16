package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateBrandPublishedVersionCmd = &cobra.Command{
	Use:   "update-brand-published-version",
	Short: "Updates the published version of a brand.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateBrandPublishedVersionCmd).Standalone()

	quicksight_updateBrandPublishedVersionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brand.")
	quicksight_updateBrandPublishedVersionCmd.Flags().String("brand-id", "", "The ID of the Quick Suite brand.")
	quicksight_updateBrandPublishedVersionCmd.Flags().String("version-id", "", "The ID of the published version.")
	quicksight_updateBrandPublishedVersionCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateBrandPublishedVersionCmd.MarkFlagRequired("brand-id")
	quicksight_updateBrandPublishedVersionCmd.MarkFlagRequired("version-id")
	quicksightCmd.AddCommand(quicksight_updateBrandPublishedVersionCmd)
}
