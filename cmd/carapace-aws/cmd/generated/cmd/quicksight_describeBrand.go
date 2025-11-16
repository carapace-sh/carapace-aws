package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeBrandCmd = &cobra.Command{
	Use:   "describe-brand",
	Short: "Describes a brand.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeBrandCmd).Standalone()

	quicksight_describeBrandCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brand.")
	quicksight_describeBrandCmd.Flags().String("brand-id", "", "The ID of the Quick Suite brand.")
	quicksight_describeBrandCmd.Flags().String("version-id", "", "The ID of the specific version.")
	quicksight_describeBrandCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeBrandCmd.MarkFlagRequired("brand-id")
	quicksightCmd.AddCommand(quicksight_describeBrandCmd)
}
