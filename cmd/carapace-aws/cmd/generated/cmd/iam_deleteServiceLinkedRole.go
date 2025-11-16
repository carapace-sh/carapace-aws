package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteServiceLinkedRoleCmd = &cobra.Command{
	Use:   "delete-service-linked-role",
	Short: "Submits a service-linked role deletion request and returns a `DeletionTaskId`, which you can use to check the status of the deletion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteServiceLinkedRoleCmd).Standalone()

	iam_deleteServiceLinkedRoleCmd.Flags().String("role-name", "", "The name of the service-linked role to be deleted.")
	iam_deleteServiceLinkedRoleCmd.MarkFlagRequired("role-name")
	iamCmd.AddCommand(iam_deleteServiceLinkedRoleCmd)
}
