package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_rejectAssignmentCmd = &cobra.Command{
	Use:   "reject-assignment",
	Short: "The `RejectAssignment` operation rejects the results of a completed assignment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_rejectAssignmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_rejectAssignmentCmd).Standalone()

		mturk_rejectAssignmentCmd.Flags().String("assignment-id", "", "The ID of the assignment.")
		mturk_rejectAssignmentCmd.Flags().String("requester-feedback", "", "A message for the Worker, which the Worker can see in the Status section of the web site.")
		mturk_rejectAssignmentCmd.MarkFlagRequired("assignment-id")
		mturk_rejectAssignmentCmd.MarkFlagRequired("requester-feedback")
	})
	mturkCmd.AddCommand(mturk_rejectAssignmentCmd)
}
