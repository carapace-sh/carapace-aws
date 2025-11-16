package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_batchGetSchemaAnalysisRuleCmd = &cobra.Command{
	Use:   "batch-get-schema-analysis-rule",
	Short: "Retrieves multiple analysis rule schemas.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_batchGetSchemaAnalysisRuleCmd).Standalone()

	cleanrooms_batchGetSchemaAnalysisRuleCmd.Flags().String("collaboration-identifier", "", "The unique identifier of the collaboration that contains the schema analysis rule.")
	cleanrooms_batchGetSchemaAnalysisRuleCmd.Flags().String("schema-analysis-rule-requests", "", "The information that's necessary to retrieve a schema analysis rule.")
	cleanrooms_batchGetSchemaAnalysisRuleCmd.MarkFlagRequired("collaboration-identifier")
	cleanrooms_batchGetSchemaAnalysisRuleCmd.MarkFlagRequired("schema-analysis-rule-requests")
	cleanroomsCmd.AddCommand(cleanrooms_batchGetSchemaAnalysisRuleCmd)
}
