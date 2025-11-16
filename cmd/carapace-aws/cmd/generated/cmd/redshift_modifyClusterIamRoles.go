package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyClusterIamRolesCmd = &cobra.Command{
	Use:   "modify-cluster-iam-roles",
	Short: "Modifies the list of Identity and Access Management (IAM) roles that can be used by the cluster to access other Amazon Web Services services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyClusterIamRolesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_modifyClusterIamRolesCmd).Standalone()

		redshift_modifyClusterIamRolesCmd.Flags().String("add-iam-roles", "", "Zero or more IAM roles to associate with the cluster.")
		redshift_modifyClusterIamRolesCmd.Flags().String("cluster-identifier", "", "The unique identifier of the cluster for which you want to associate or disassociate IAM roles.")
		redshift_modifyClusterIamRolesCmd.Flags().String("default-iam-role-arn", "", "The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was last modified.")
		redshift_modifyClusterIamRolesCmd.Flags().String("remove-iam-roles", "", "Zero or more IAM roles in ARN format to disassociate from the cluster.")
		redshift_modifyClusterIamRolesCmd.MarkFlagRequired("cluster-identifier")
	})
	redshiftCmd.AddCommand(redshift_modifyClusterIamRolesCmd)
}
