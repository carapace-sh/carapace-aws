package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteQueryDefinitionCmd = &cobra.Command{
	Use:   "delete-query-definition",
	Short: "Deletes a saved CloudWatch Logs Insights query definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteQueryDefinitionCmd).Standalone()

	logs_deleteQueryDefinitionCmd.Flags().String("query-definition-id", "", "The ID of the query definition that you want to delete.")
	logs_deleteQueryDefinitionCmd.MarkFlagRequired("query-definition-id")
	logsCmd.AddCommand(logs_deleteQueryDefinitionCmd)
}
