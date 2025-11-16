package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getIntegrationTablePropertiesCmd = &cobra.Command{
	Use:   "get-integration-table-properties",
	Short: "This API is used to retrieve optional override properties for the tables that need to be replicated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getIntegrationTablePropertiesCmd).Standalone()

	glue_getIntegrationTablePropertiesCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the target table for which to retrieve integration table properties.")
	glue_getIntegrationTablePropertiesCmd.Flags().String("table-name", "", "The name of the table to be replicated.")
	glue_getIntegrationTablePropertiesCmd.MarkFlagRequired("resource-arn")
	glue_getIntegrationTablePropertiesCmd.MarkFlagRequired("table-name")
	glueCmd.AddCommand(glue_getIntegrationTablePropertiesCmd)
}
