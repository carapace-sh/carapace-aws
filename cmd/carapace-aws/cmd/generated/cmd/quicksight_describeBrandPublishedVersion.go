package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeBrandPublishedVersionCmd = &cobra.Command{
	Use:   "describe-brand-published-version",
	Short: "Describes the published version of the brand.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeBrandPublishedVersionCmd).Standalone()

	quicksight_describeBrandPublishedVersionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brand.")
	quicksight_describeBrandPublishedVersionCmd.Flags().String("brand-id", "", "The ID of the Quick Suite brand.")
	quicksight_describeBrandPublishedVersionCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeBrandPublishedVersionCmd.MarkFlagRequired("brand-id")
	quicksightCmd.AddCommand(quicksight_describeBrandPublishedVersionCmd)
}
