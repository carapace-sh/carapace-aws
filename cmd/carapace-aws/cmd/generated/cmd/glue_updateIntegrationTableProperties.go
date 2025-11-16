package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateIntegrationTablePropertiesCmd = &cobra.Command{
	Use:   "update-integration-table-properties",
	Short: "This API is used to provide optional override properties for the tables that need to be replicated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateIntegrationTablePropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_updateIntegrationTablePropertiesCmd).Standalone()

		glue_updateIntegrationTablePropertiesCmd.Flags().String("resource-arn", "", "The connection ARN of the source, or the database ARN of the target.")
		glue_updateIntegrationTablePropertiesCmd.Flags().String("source-table-config", "", "A structure for the source table configuration.")
		glue_updateIntegrationTablePropertiesCmd.Flags().String("table-name", "", "The name of the table to be replicated.")
		glue_updateIntegrationTablePropertiesCmd.Flags().String("target-table-config", "", "A structure for the target table configuration.")
		glue_updateIntegrationTablePropertiesCmd.MarkFlagRequired("resource-arn")
		glue_updateIntegrationTablePropertiesCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_updateIntegrationTablePropertiesCmd)
}
