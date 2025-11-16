package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_removeRoleFromDbinstanceCmd = &cobra.Command{
	Use:   "remove-role-from-dbinstance",
	Short: "Disassociates an Amazon Web Services Identity and Access Management (IAM) role from a DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_removeRoleFromDbinstanceCmd).Standalone()

	rds_removeRoleFromDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The name of the DB instance to disassociate the IAM role from.")
	rds_removeRoleFromDbinstanceCmd.Flags().String("feature-name", "", "The name of the feature for the DB instance that the IAM role is to be disassociated from.")
	rds_removeRoleFromDbinstanceCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to disassociate from the DB instance, for example, `arn:aws:iam::123456789012:role/AccessRole`.")
	rds_removeRoleFromDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	rds_removeRoleFromDbinstanceCmd.MarkFlagRequired("feature-name")
	rds_removeRoleFromDbinstanceCmd.MarkFlagRequired("role-arn")
	rdsCmd.AddCommand(rds_removeRoleFromDbinstanceCmd)
}
