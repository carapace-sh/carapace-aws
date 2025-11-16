package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeQueryDefinitionsCmd = &cobra.Command{
	Use:   "describe-query-definitions",
	Short: "This operation returns a paginated list of your saved CloudWatch Logs Insights query definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeQueryDefinitionsCmd).Standalone()

	logs_describeQueryDefinitionsCmd.Flags().String("max-results", "", "Limits the number of returned query definitions to the specified number.")
	logs_describeQueryDefinitionsCmd.Flags().String("next-token", "", "")
	logs_describeQueryDefinitionsCmd.Flags().String("query-definition-name-prefix", "", "Use this parameter to filter your results to only the query definitions that have names that start with the prefix you specify.")
	logs_describeQueryDefinitionsCmd.Flags().String("query-language", "", "The query language used for this query.")
	logsCmd.AddCommand(logs_describeQueryDefinitionsCmd)
}
