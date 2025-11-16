package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getServiceLinkedRoleDeletionStatusCmd = &cobra.Command{
	Use:   "get-service-linked-role-deletion-status",
	Short: "Retrieves the status of your service-linked role deletion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getServiceLinkedRoleDeletionStatusCmd).Standalone()

	iam_getServiceLinkedRoleDeletionStatusCmd.Flags().String("deletion-task-id", "", "The deletion task identifier.")
	iam_getServiceLinkedRoleDeletionStatusCmd.MarkFlagRequired("deletion-task-id")
	iamCmd.AddCommand(iam_getServiceLinkedRoleDeletionStatusCmd)
}
