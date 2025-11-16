package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_approveAssignmentCmd = &cobra.Command{
	Use:   "approve-assignment",
	Short: "The `ApproveAssignment` operation approves the results of a completed assignment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_approveAssignmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_approveAssignmentCmd).Standalone()

		mturk_approveAssignmentCmd.Flags().String("assignment-id", "", "The ID of the assignment.")
		mturk_approveAssignmentCmd.Flags().Bool("no-override-rejection", false, "A flag indicating that an assignment should be approved even if it was previously rejected.")
		mturk_approveAssignmentCmd.Flags().Bool("override-rejection", false, "A flag indicating that an assignment should be approved even if it was previously rejected.")
		mturk_approveAssignmentCmd.Flags().String("requester-feedback", "", "A message for the Worker, which the Worker can see in the Status section of the web site.")
		mturk_approveAssignmentCmd.MarkFlagRequired("assignment-id")
		mturk_approveAssignmentCmd.Flag("no-override-rejection").Hidden = true
	})
	mturkCmd.AddCommand(mturk_approveAssignmentCmd)
}
