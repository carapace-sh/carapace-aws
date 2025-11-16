package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_restoreTableFromClusterSnapshotCmd = &cobra.Command{
	Use:   "restore-table-from-cluster-snapshot",
	Short: "Creates a new table from a table in an Amazon Redshift cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_restoreTableFromClusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_restoreTableFromClusterSnapshotCmd).Standalone()

		redshift_restoreTableFromClusterSnapshotCmd.Flags().String("cluster-identifier", "", "The identifier of the Amazon Redshift cluster to restore the table to.")
		redshift_restoreTableFromClusterSnapshotCmd.Flags().String("enable-case-sensitive-identifier", "", "Indicates whether name identifiers for database, schema, and table are case sensitive.")
		redshift_restoreTableFromClusterSnapshotCmd.Flags().String("new-table-name", "", "The name of the table to create as a result of the current request.")
		redshift_restoreTableFromClusterSnapshotCmd.Flags().String("snapshot-identifier", "", "The identifier of the snapshot to restore the table from.")
		redshift_restoreTableFromClusterSnapshotCmd.Flags().String("source-database-name", "", "The name of the source database that contains the table to restore from.")
		redshift_restoreTableFromClusterSnapshotCmd.Flags().String("source-schema-name", "", "The name of the source schema that contains the table to restore from.")
		redshift_restoreTableFromClusterSnapshotCmd.Flags().String("source-table-name", "", "The name of the source table to restore from.")
		redshift_restoreTableFromClusterSnapshotCmd.Flags().String("target-database-name", "", "The name of the database to restore the table to.")
		redshift_restoreTableFromClusterSnapshotCmd.Flags().String("target-schema-name", "", "The name of the schema to restore the table to.")
		redshift_restoreTableFromClusterSnapshotCmd.MarkFlagRequired("cluster-identifier")
		redshift_restoreTableFromClusterSnapshotCmd.MarkFlagRequired("new-table-name")
		redshift_restoreTableFromClusterSnapshotCmd.MarkFlagRequired("snapshot-identifier")
		redshift_restoreTableFromClusterSnapshotCmd.MarkFlagRequired("source-database-name")
		redshift_restoreTableFromClusterSnapshotCmd.MarkFlagRequired("source-table-name")
	})
	redshiftCmd.AddCommand(redshift_restoreTableFromClusterSnapshotCmd)
}
