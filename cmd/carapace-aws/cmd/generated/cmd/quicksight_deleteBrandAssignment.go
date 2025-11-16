package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteBrandAssignmentCmd = &cobra.Command{
	Use:   "delete-brand-assignment",
	Short: "Deletes a brand assignment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteBrandAssignmentCmd).Standalone()

	quicksight_deleteBrandAssignmentCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brand assignment.")
	quicksight_deleteBrandAssignmentCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_deleteBrandAssignmentCmd)
}
