package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyTenantDatabaseCmd = &cobra.Command{
	Use:   "modify-tenant-database",
	Short: "Modifies an existing tenant database in a DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyTenantDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_modifyTenantDatabaseCmd).Standalone()

		rds_modifyTenantDatabaseCmd.Flags().String("dbinstance-identifier", "", "The identifier of the DB instance that contains the tenant database that you are modifying.")
		rds_modifyTenantDatabaseCmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.")
		rds_modifyTenantDatabaseCmd.Flags().String("master-user-password", "", "The new password for the master user of the specified tenant database in your DB instance.")
		rds_modifyTenantDatabaseCmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
		rds_modifyTenantDatabaseCmd.Flags().String("new-tenant-dbname", "", "The new name of the tenant database when renaming a tenant database.")
		rds_modifyTenantDatabaseCmd.Flags().String("rotate-master-user-password", "", "Specifies whether to rotate the secret managed by Amazon Web Services Secrets Manager for the master user password.")
		rds_modifyTenantDatabaseCmd.Flags().String("tenant-dbname", "", "The user-supplied name of the tenant database that you want to modify.")
		rds_modifyTenantDatabaseCmd.MarkFlagRequired("dbinstance-identifier")
		rds_modifyTenantDatabaseCmd.MarkFlagRequired("tenant-dbname")
	})
	rdsCmd.AddCommand(rds_modifyTenantDatabaseCmd)
}
