package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_removeRoleFromDbclusterCmd = &cobra.Command{
	Use:   "remove-role-from-dbcluster",
	Short: "Removes the asssociation of an Amazon Web Services Identity and Access Management (IAM) role from a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_removeRoleFromDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_removeRoleFromDbclusterCmd).Standalone()

		rds_removeRoleFromDbclusterCmd.Flags().String("dbcluster-identifier", "", "The name of the DB cluster to disassociate the IAM role from.")
		rds_removeRoleFromDbclusterCmd.Flags().String("feature-name", "", "The name of the feature for the DB cluster that the IAM role is to be disassociated from.")
		rds_removeRoleFromDbclusterCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to disassociate from the Aurora DB cluster, for example `arn:aws:iam::123456789012:role/AuroraAccessRole`.")
		rds_removeRoleFromDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
		rds_removeRoleFromDbclusterCmd.MarkFlagRequired("role-arn")
	})
	rdsCmd.AddCommand(rds_removeRoleFromDbclusterCmd)
}
