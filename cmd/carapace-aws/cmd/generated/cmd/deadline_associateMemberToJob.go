package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_associateMemberToJobCmd = &cobra.Command{
	Use:   "associate-member-to-job",
	Short: "Assigns a job membership level to a member",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_associateMemberToJobCmd).Standalone()

	deadline_associateMemberToJobCmd.Flags().String("farm-id", "", "The farm ID of the job to associate with the member.")
	deadline_associateMemberToJobCmd.Flags().String("identity-store-id", "", "The member's identity store ID to associate with the job.")
	deadline_associateMemberToJobCmd.Flags().String("job-id", "", "The job ID to associate with the member.")
	deadline_associateMemberToJobCmd.Flags().String("membership-level", "", "The principal's membership level for the associated job.")
	deadline_associateMemberToJobCmd.Flags().String("principal-id", "", "The member's principal ID to associate with the job.")
	deadline_associateMemberToJobCmd.Flags().String("principal-type", "", "The member's principal type to associate with the job.")
	deadline_associateMemberToJobCmd.Flags().String("queue-id", "", "The queue ID to associate to the member.")
	deadline_associateMemberToJobCmd.MarkFlagRequired("farm-id")
	deadline_associateMemberToJobCmd.MarkFlagRequired("identity-store-id")
	deadline_associateMemberToJobCmd.MarkFlagRequired("job-id")
	deadline_associateMemberToJobCmd.MarkFlagRequired("membership-level")
	deadline_associateMemberToJobCmd.MarkFlagRequired("principal-id")
	deadline_associateMemberToJobCmd.MarkFlagRequired("principal-type")
	deadline_associateMemberToJobCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_associateMemberToJobCmd)
}
