package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbclusterAutomatedBackupCmd = &cobra.Command{
	Use:   "delete-dbcluster-automated-backup",
	Short: "Deletes automated backups using the `DbClusterResourceId` value of the source DB cluster or the Amazon Resource Name (ARN) of the automated backups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbclusterAutomatedBackupCmd).Standalone()

	rds_deleteDbclusterAutomatedBackupCmd.Flags().String("db-cluster-resource-id", "", "The identifier for the source DB cluster, which can't be changed and which is unique to an Amazon Web Services Region.")
	rds_deleteDbclusterAutomatedBackupCmd.MarkFlagRequired("db-cluster-resource-id")
	rdsCmd.AddCommand(rds_deleteDbclusterAutomatedBackupCmd)
}
