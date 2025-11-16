package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getSchemaAnalysisRuleCmd = &cobra.Command{
	Use:   "get-schema-analysis-rule",
	Short: "Retrieves a schema analysis rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getSchemaAnalysisRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getSchemaAnalysisRuleCmd).Standalone()

		cleanrooms_getSchemaAnalysisRuleCmd.Flags().String("collaboration-identifier", "", "A unique identifier for the collaboration that the schema belongs to.")
		cleanrooms_getSchemaAnalysisRuleCmd.Flags().String("name", "", "The name of the schema to retrieve the analysis rule for.")
		cleanrooms_getSchemaAnalysisRuleCmd.Flags().String("type", "", "The type of the schema analysis rule to retrieve.")
		cleanrooms_getSchemaAnalysisRuleCmd.MarkFlagRequired("collaboration-identifier")
		cleanrooms_getSchemaAnalysisRuleCmd.MarkFlagRequired("name")
		cleanrooms_getSchemaAnalysisRuleCmd.MarkFlagRequired("type")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getSchemaAnalysisRuleCmd)
}
