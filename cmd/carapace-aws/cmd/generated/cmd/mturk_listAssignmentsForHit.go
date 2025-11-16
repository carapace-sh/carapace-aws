package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listAssignmentsForHitCmd = &cobra.Command{
	Use:   "list-assignments-for-hit",
	Short: "The `ListAssignmentsForHIT` operation retrieves completed assignments for a HIT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listAssignmentsForHitCmd).Standalone()

	mturk_listAssignmentsForHitCmd.Flags().String("assignment-statuses", "", "The status of the assignments to return: Submitted | Approved | Rejected")
	mturk_listAssignmentsForHitCmd.Flags().String("hitid", "", "The ID of the HIT.")
	mturk_listAssignmentsForHitCmd.Flags().String("max-results", "", "")
	mturk_listAssignmentsForHitCmd.Flags().String("next-token", "", "Pagination token")
	mturk_listAssignmentsForHitCmd.MarkFlagRequired("hitid")
	mturkCmd.AddCommand(mturk_listAssignmentsForHitCmd)
}
