package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeBrandAssignmentCmd = &cobra.Command{
	Use:   "describe-brand-assignment",
	Short: "Describes a brand assignment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeBrandAssignmentCmd).Standalone()

	quicksight_describeBrandAssignmentCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brand assignment.")
	quicksight_describeBrandAssignmentCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeBrandAssignmentCmd)
}
