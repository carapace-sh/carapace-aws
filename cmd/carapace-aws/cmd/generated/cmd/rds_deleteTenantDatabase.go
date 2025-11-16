package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteTenantDatabaseCmd = &cobra.Command{
	Use:   "delete-tenant-database",
	Short: "Deletes a tenant database from your DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteTenantDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteTenantDatabaseCmd).Standalone()

		rds_deleteTenantDatabaseCmd.Flags().String("dbinstance-identifier", "", "The user-supplied identifier for the DB instance that contains the tenant database that you want to delete.")
		rds_deleteTenantDatabaseCmd.Flags().String("final-dbsnapshot-identifier", "", "The `DBSnapshotIdentifier` of the new `DBSnapshot` created when the `SkipFinalSnapshot` parameter is disabled.")
		rds_deleteTenantDatabaseCmd.Flags().Bool("no-skip-final-snapshot", false, "Specifies whether to skip the creation of a final DB snapshot before removing the tenant database from your DB instance.")
		rds_deleteTenantDatabaseCmd.Flags().Bool("skip-final-snapshot", false, "Specifies whether to skip the creation of a final DB snapshot before removing the tenant database from your DB instance.")
		rds_deleteTenantDatabaseCmd.Flags().String("tenant-dbname", "", "The user-supplied name of the tenant database that you want to remove from your DB instance.")
		rds_deleteTenantDatabaseCmd.MarkFlagRequired("dbinstance-identifier")
		rds_deleteTenantDatabaseCmd.Flag("no-skip-final-snapshot").Hidden = true
		rds_deleteTenantDatabaseCmd.MarkFlagRequired("tenant-dbname")
	})
	rdsCmd.AddCommand(rds_deleteTenantDatabaseCmd)
}
