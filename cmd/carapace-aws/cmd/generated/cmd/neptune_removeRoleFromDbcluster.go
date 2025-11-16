package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_removeRoleFromDbclusterCmd = &cobra.Command{
	Use:   "remove-role-from-dbcluster",
	Short: "Disassociates an Identity and Access Management (IAM) role from a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_removeRoleFromDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_removeRoleFromDbclusterCmd).Standalone()

		neptune_removeRoleFromDbclusterCmd.Flags().String("dbcluster-identifier", "", "The name of the DB cluster to disassociate the IAM role from.")
		neptune_removeRoleFromDbclusterCmd.Flags().String("feature-name", "", "The name of the feature for the DB cluster that the IAM role is to be disassociated from.")
		neptune_removeRoleFromDbclusterCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to disassociate from the DB cluster, for example `arn:aws:iam::123456789012:role/NeptuneAccessRole`.")
		neptune_removeRoleFromDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
		neptune_removeRoleFromDbclusterCmd.MarkFlagRequired("role-arn")
	})
	neptuneCmd.AddCommand(neptune_removeRoleFromDbclusterCmd)
}
