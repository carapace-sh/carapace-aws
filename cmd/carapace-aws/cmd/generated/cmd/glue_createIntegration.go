package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createIntegrationCmd = &cobra.Command{
	Use:   "create-integration",
	Short: "Creates a Zero-ETL integration in the caller's account between two resources with Amazon Resource Names (ARNs): the `SourceArn` and `TargetArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createIntegrationCmd).Standalone()

	glue_createIntegrationCmd.Flags().String("additional-encryption-context", "", "An optional set of non-secret key–value pairs that contains additional contextual information for encryption.")
	glue_createIntegrationCmd.Flags().String("data-filter", "", "Selects source tables for the integration using Maxwell filter syntax.")
	glue_createIntegrationCmd.Flags().String("description", "", "A description of the integration.")
	glue_createIntegrationCmd.Flags().String("integration-config", "", "The configuration settings.")
	glue_createIntegrationCmd.Flags().String("integration-name", "", "A unique name for an integration in Glue.")
	glue_createIntegrationCmd.Flags().String("kms-key-id", "", "The ARN of a KMS key used for encrypting the channel.")
	glue_createIntegrationCmd.Flags().String("source-arn", "", "The ARN of the source resource for the integration.")
	glue_createIntegrationCmd.Flags().String("tags", "", "Metadata assigned to the resource consisting of a list of key-value pairs.")
	glue_createIntegrationCmd.Flags().String("target-arn", "", "The ARN of the target resource for the integration.")
	glue_createIntegrationCmd.MarkFlagRequired("integration-name")
	glue_createIntegrationCmd.MarkFlagRequired("source-arn")
	glue_createIntegrationCmd.MarkFlagRequired("target-arn")
	glueCmd.AddCommand(glue_createIntegrationCmd)
}
