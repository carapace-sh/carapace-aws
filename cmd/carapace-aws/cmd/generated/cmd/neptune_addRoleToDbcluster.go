package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_addRoleToDbclusterCmd = &cobra.Command{
	Use:   "add-role-to-dbcluster",
	Short: "Associates an Identity and Access Management (IAM) role with an Neptune DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_addRoleToDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_addRoleToDbclusterCmd).Standalone()

		neptune_addRoleToDbclusterCmd.Flags().String("dbcluster-identifier", "", "The name of the DB cluster to associate the IAM role with.")
		neptune_addRoleToDbclusterCmd.Flags().String("feature-name", "", "The name of the feature for the Neptune DB cluster that the IAM role is to be associated with.")
		neptune_addRoleToDbclusterCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to associate with the Neptune DB cluster, for example `arn:aws:iam::123456789012:role/NeptuneAccessRole`.")
		neptune_addRoleToDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
		neptune_addRoleToDbclusterCmd.MarkFlagRequired("role-arn")
	})
	neptuneCmd.AddCommand(neptune_addRoleToDbclusterCmd)
}
