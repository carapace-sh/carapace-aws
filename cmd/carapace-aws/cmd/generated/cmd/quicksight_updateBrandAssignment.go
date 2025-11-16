package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateBrandAssignmentCmd = &cobra.Command{
	Use:   "update-brand-assignment",
	Short: "Updates a brand assignment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateBrandAssignmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateBrandAssignmentCmd).Standalone()

		quicksight_updateBrandAssignmentCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that owns the brand assignment.")
		quicksight_updateBrandAssignmentCmd.Flags().String("brand-arn", "", "The Amazon Resource Name (ARN) of the brand.")
		quicksight_updateBrandAssignmentCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateBrandAssignmentCmd.MarkFlagRequired("brand-arn")
	})
	quicksightCmd.AddCommand(quicksight_updateBrandAssignmentCmd)
}
