package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createIntegrationCmd = &cobra.Command{
	Use:   "create-integration",
	Short: "Creates a zero-ETL integration or S3 event integration with Amazon Redshift.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createIntegrationCmd).Standalone()

	redshift_createIntegrationCmd.Flags().String("additional-encryption-context", "", "An optional set of non-secret key–value pairs that contains additional contextual information about the data.")
	redshift_createIntegrationCmd.Flags().String("description", "", "A description of the integration.")
	redshift_createIntegrationCmd.Flags().String("integration-name", "", "The name of the integration.")
	redshift_createIntegrationCmd.Flags().String("kmskey-id", "", "An Key Management Service (KMS) key identifier for the key to use to encrypt the integration.")
	redshift_createIntegrationCmd.Flags().String("source-arn", "", "The Amazon Resource Name (ARN) of the database to use as the source for replication.")
	redshift_createIntegrationCmd.Flags().String("tag-list", "", "A list of tags.")
	redshift_createIntegrationCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) of the Amazon Redshift data warehouse to use as the target for replication.")
	redshift_createIntegrationCmd.MarkFlagRequired("integration-name")
	redshift_createIntegrationCmd.MarkFlagRequired("source-arn")
	redshift_createIntegrationCmd.MarkFlagRequired("target-arn")
	redshiftCmd.AddCommand(redshift_createIntegrationCmd)
}
