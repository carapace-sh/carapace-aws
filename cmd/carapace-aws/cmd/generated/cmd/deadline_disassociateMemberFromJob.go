package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_disassociateMemberFromJobCmd = &cobra.Command{
	Use:   "disassociate-member-from-job",
	Short: "Disassociates a member from a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_disassociateMemberFromJobCmd).Standalone()

	deadline_disassociateMemberFromJobCmd.Flags().String("farm-id", "", "The farm ID for the job to disassociate from the member.")
	deadline_disassociateMemberFromJobCmd.Flags().String("job-id", "", "The job ID to disassociate from a member in a job.")
	deadline_disassociateMemberFromJobCmd.Flags().String("principal-id", "", "A member's principal ID to disassociate from a job.")
	deadline_disassociateMemberFromJobCmd.Flags().String("queue-id", "", "The queue ID connected to a job for which you're disassociating a member.")
	deadline_disassociateMemberFromJobCmd.MarkFlagRequired("farm-id")
	deadline_disassociateMemberFromJobCmd.MarkFlagRequired("job-id")
	deadline_disassociateMemberFromJobCmd.MarkFlagRequired("principal-id")
	deadline_disassociateMemberFromJobCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_disassociateMemberFromJobCmd)
}
