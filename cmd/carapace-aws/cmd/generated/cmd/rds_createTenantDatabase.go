package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createTenantDatabaseCmd = &cobra.Command{
	Use:   "create-tenant-database",
	Short: "Creates a tenant database in a DB instance that uses the multi-tenant configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createTenantDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_createTenantDatabaseCmd).Standalone()

		rds_createTenantDatabaseCmd.Flags().String("character-set-name", "", "The character set for your tenant database.")
		rds_createTenantDatabaseCmd.Flags().String("dbinstance-identifier", "", "The user-supplied DB instance identifier.")
		rds_createTenantDatabaseCmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.")
		rds_createTenantDatabaseCmd.Flags().String("master-user-password", "", "The password for the master user in your tenant database.")
		rds_createTenantDatabaseCmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
		rds_createTenantDatabaseCmd.Flags().String("master-username", "", "The name for the master user account in your tenant database.")
		rds_createTenantDatabaseCmd.Flags().String("nchar-character-set-name", "", "The `NCHAR` value for the tenant database.")
		rds_createTenantDatabaseCmd.Flags().String("tags", "", "")
		rds_createTenantDatabaseCmd.Flags().String("tenant-dbname", "", "The user-supplied name of the tenant database that you want to create in your DB instance.")
		rds_createTenantDatabaseCmd.MarkFlagRequired("dbinstance-identifier")
		rds_createTenantDatabaseCmd.MarkFlagRequired("master-username")
		rds_createTenantDatabaseCmd.MarkFlagRequired("tenant-dbname")
	})
	rdsCmd.AddCommand(rds_createTenantDatabaseCmd)
}
