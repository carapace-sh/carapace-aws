package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteIntegrationTablePropertiesCmd = &cobra.Command{
	Use:   "delete-integration-table-properties",
	Short: "Deletes the table properties that have been created for the tables that need to be replicated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteIntegrationTablePropertiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteIntegrationTablePropertiesCmd).Standalone()

		glue_deleteIntegrationTablePropertiesCmd.Flags().String("resource-arn", "", "The connection ARN of the source, or the database ARN of the target.")
		glue_deleteIntegrationTablePropertiesCmd.Flags().String("table-name", "", "The name of the table to be replicated.")
		glue_deleteIntegrationTablePropertiesCmd.MarkFlagRequired("resource-arn")
		glue_deleteIntegrationTablePropertiesCmd.MarkFlagRequired("table-name")
	})
	glueCmd.AddCommand(glue_deleteIntegrationTablePropertiesCmd)
}
