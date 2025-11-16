package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_restoreTableFromSnapshotCmd = &cobra.Command{
	Use:   "restore-table-from-snapshot",
	Short: "Restores a table from a snapshot to your Amazon Redshift Serverless instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_restoreTableFromSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_restoreTableFromSnapshotCmd).Standalone()

		redshiftServerless_restoreTableFromSnapshotCmd.Flags().Bool("activate-case-sensitive-identifier", false, "Indicates whether name identifiers for database, schema, and table are case sensitive.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().String("namespace-name", "", "The namespace of the snapshot to restore from.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().String("new-table-name", "", "The name of the table to create from the restore operation.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().Bool("no-activate-case-sensitive-identifier", false, "Indicates whether name identifiers for database, schema, and table are case sensitive.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().String("snapshot-name", "", "The name of the snapshot to restore the table from.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().String("source-database-name", "", "The name of the source database that contains the table being restored.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().String("source-schema-name", "", "The name of the source schema that contains the table being restored.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().String("source-table-name", "", "The name of the source table being restored.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().String("target-database-name", "", "The name of the database to restore the table to.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().String("target-schema-name", "", "The name of the schema to restore the table to.")
		redshiftServerless_restoreTableFromSnapshotCmd.Flags().String("workgroup-name", "", "The workgroup to restore the table to.")
		redshiftServerless_restoreTableFromSnapshotCmd.MarkFlagRequired("namespace-name")
		redshiftServerless_restoreTableFromSnapshotCmd.MarkFlagRequired("new-table-name")
		redshiftServerless_restoreTableFromSnapshotCmd.Flag("no-activate-case-sensitive-identifier").Hidden = true
		redshiftServerless_restoreTableFromSnapshotCmd.MarkFlagRequired("snapshot-name")
		redshiftServerless_restoreTableFromSnapshotCmd.MarkFlagRequired("source-database-name")
		redshiftServerless_restoreTableFromSnapshotCmd.MarkFlagRequired("source-table-name")
		redshiftServerless_restoreTableFromSnapshotCmd.MarkFlagRequired("workgroup-name")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_restoreTableFromSnapshotCmd)
}
