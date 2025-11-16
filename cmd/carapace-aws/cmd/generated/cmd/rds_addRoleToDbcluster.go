package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_addRoleToDbclusterCmd = &cobra.Command{
	Use:   "add-role-to-dbcluster",
	Short: "Associates an Identity and Access Management (IAM) role with a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_addRoleToDbclusterCmd).Standalone()

	rds_addRoleToDbclusterCmd.Flags().String("dbcluster-identifier", "", "The name of the DB cluster to associate the IAM role with.")
	rds_addRoleToDbclusterCmd.Flags().String("feature-name", "", "The name of the feature for the DB cluster that the IAM role is to be associated with.")
	rds_addRoleToDbclusterCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to associate with the Aurora DB cluster, for example `arn:aws:iam::123456789012:role/AuroraAccessRole`.")
	rds_addRoleToDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	rds_addRoleToDbclusterCmd.MarkFlagRequired("role-arn")
	rdsCmd.AddCommand(rds_addRoleToDbclusterCmd)
}
