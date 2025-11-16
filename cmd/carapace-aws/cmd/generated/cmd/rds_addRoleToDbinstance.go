package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_addRoleToDbinstanceCmd = &cobra.Command{
	Use:   "add-role-to-dbinstance",
	Short: "Associates an Amazon Web Services Identity and Access Management (IAM) role with a DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_addRoleToDbinstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_addRoleToDbinstanceCmd).Standalone()

		rds_addRoleToDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The name of the DB instance to associate the IAM role with.")
		rds_addRoleToDbinstanceCmd.Flags().String("feature-name", "", "The name of the feature for the DB instance that the IAM role is to be associated with.")
		rds_addRoleToDbinstanceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to associate with the DB instance, for example `arn:aws:iam::123456789012:role/AccessRole`.")
		rds_addRoleToDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
		rds_addRoleToDbinstanceCmd.MarkFlagRequired("feature-name")
		rds_addRoleToDbinstanceCmd.MarkFlagRequired("role-arn")
	})
	rdsCmd.AddCommand(rds_addRoleToDbinstanceCmd)
}
