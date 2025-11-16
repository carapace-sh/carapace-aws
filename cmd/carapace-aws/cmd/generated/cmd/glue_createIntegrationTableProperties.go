package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createIntegrationTablePropertiesCmd = &cobra.Command{
	Use:   "create-integration-table-properties",
	Short: "This API is used to provide optional override properties for the the tables that need to be replicated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createIntegrationTablePropertiesCmd).Standalone()

	glue_createIntegrationTablePropertiesCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the target table for which to create integration table properties.")
	glue_createIntegrationTablePropertiesCmd.Flags().String("source-table-config", "", "A structure for the source table configuration.")
	glue_createIntegrationTablePropertiesCmd.Flags().String("table-name", "", "The name of the table to be replicated.")
	glue_createIntegrationTablePropertiesCmd.Flags().String("target-table-config", "", "A structure for the target table configuration.")
	glue_createIntegrationTablePropertiesCmd.MarkFlagRequired("resource-arn")
	glue_createIntegrationTablePropertiesCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_createIntegrationTablePropertiesCmd)
}
