package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeTenantDatabasesCmd = &cobra.Command{
	Use:   "describe-tenant-databases",
	Short: "Describes the tenant databases in a DB instance that uses the multi-tenant configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeTenantDatabasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeTenantDatabasesCmd).Standalone()

		rds_describeTenantDatabasesCmd.Flags().String("dbinstance-identifier", "", "The user-supplied DB instance identifier, which must match the identifier of an existing instance owned by the Amazon Web Services account.")
		rds_describeTenantDatabasesCmd.Flags().String("filters", "", "A filter that specifies one or more database tenants to describe.")
		rds_describeTenantDatabasesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeTenantDatabases` request.")
		rds_describeTenantDatabasesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeTenantDatabasesCmd.Flags().String("tenant-dbname", "", "The user-supplied tenant database name, which must match the name of an existing tenant database on the specified DB instance owned by your Amazon Web Services account.")
	})
	rdsCmd.AddCommand(rds_describeTenantDatabasesCmd)
}
