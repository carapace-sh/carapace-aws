package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createIntegrationCmd = &cobra.Command{
	Use:   "create-integration",
	Short: "Creates a zero-ETL integration with Amazon Redshift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_createIntegrationCmd).Standalone()

		rds_createIntegrationCmd.Flags().String("additional-encryption-context", "", "An optional set of non-secret key–value pairs that contains additional contextual information about the data.")
		rds_createIntegrationCmd.Flags().String("data-filter", "", "Data filtering options for the integration.")
		rds_createIntegrationCmd.Flags().String("description", "", "A description of the integration.")
		rds_createIntegrationCmd.Flags().String("integration-name", "", "The name of the integration.")
		rds_createIntegrationCmd.Flags().String("kmskey-id", "", "The Amazon Web Services Key Management System (Amazon Web Services KMS) key identifier for the key to use to encrypt the integration.")
		rds_createIntegrationCmd.Flags().String("source-arn", "", "The Amazon Resource Name (ARN) of the database to use as the source for replication.")
		rds_createIntegrationCmd.Flags().String("tags", "", "")
		rds_createIntegrationCmd.Flags().String("target-arn", "", "The ARN of the Redshift data warehouse to use as the target for replication.")
		rds_createIntegrationCmd.MarkFlagRequired("integration-name")
		rds_createIntegrationCmd.MarkFlagRequired("source-arn")
		rds_createIntegrationCmd.MarkFlagRequired("target-arn")
	})
	rdsCmd.AddCommand(rds_createIntegrationCmd)
}
