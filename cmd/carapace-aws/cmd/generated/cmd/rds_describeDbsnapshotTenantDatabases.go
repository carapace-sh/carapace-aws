package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbsnapshotTenantDatabasesCmd = &cobra.Command{
	Use:   "describe-dbsnapshot-tenant-databases",
	Short: "Describes the tenant databases that exist in a DB snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbsnapshotTenantDatabasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbsnapshotTenantDatabasesCmd).Standalone()

		rds_describeDbsnapshotTenantDatabasesCmd.Flags().String("dbi-resource-id", "", "A specific DB resource identifier to describe.")
		rds_describeDbsnapshotTenantDatabasesCmd.Flags().String("dbinstance-identifier", "", "The ID of the DB instance used to create the DB snapshots.")
		rds_describeDbsnapshotTenantDatabasesCmd.Flags().String("dbsnapshot-identifier", "", "The ID of a DB snapshot that contains the tenant databases to describe.")
		rds_describeDbsnapshotTenantDatabasesCmd.Flags().String("filters", "", "A filter that specifies one or more tenant databases to describe.")
		rds_describeDbsnapshotTenantDatabasesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBSnapshotTenantDatabases` request.")
		rds_describeDbsnapshotTenantDatabasesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeDbsnapshotTenantDatabasesCmd.Flags().String("snapshot-type", "", "The type of DB snapshots to be returned.")
	})
	rdsCmd.AddCommand(rds_describeDbsnapshotTenantDatabasesCmd)
}
