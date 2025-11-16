package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateManagedInstanceRoleCmd = &cobra.Command{
	Use:   "update-managed-instance-role",
	Short: "Changes the Identity and Access Management (IAM) role that is assigned to the on-premises server, edge device, or virtual machines (VM).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateManagedInstanceRoleCmd).Standalone()

	ssm_updateManagedInstanceRoleCmd.Flags().String("iam-role", "", "The name of the Identity and Access Management (IAM) role that you want to assign to the managed node.")
	ssm_updateManagedInstanceRoleCmd.Flags().String("instance-id", "", "The ID of the managed node where you want to update the role.")
	ssm_updateManagedInstanceRoleCmd.MarkFlagRequired("iam-role")
	ssm_updateManagedInstanceRoleCmd.MarkFlagRequired("instance-id")
	ssmCmd.AddCommand(ssm_updateManagedInstanceRoleCmd)
}
