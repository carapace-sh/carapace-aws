package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putQueryDefinitionCmd = &cobra.Command{
	Use:   "put-query-definition",
	Short: "Creates or updates a query definition for CloudWatch Logs Insights.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putQueryDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putQueryDefinitionCmd).Standalone()

		logs_putQueryDefinitionCmd.Flags().String("client-token", "", "Used as an idempotency token, to avoid returning an exception if the service receives the same request twice because of a network error.")
		logs_putQueryDefinitionCmd.Flags().String("log-group-names", "", "Use this parameter to include specific log groups as part of your query definition.")
		logs_putQueryDefinitionCmd.Flags().String("name", "", "A name for the query definition.")
		logs_putQueryDefinitionCmd.Flags().String("query-definition-id", "", "If you are updating a query definition, use this parameter to specify the ID of the query definition that you want to update.")
		logs_putQueryDefinitionCmd.Flags().String("query-language", "", "Specify the query language to use for this query.")
		logs_putQueryDefinitionCmd.Flags().String("query-string", "", "The query string to use for this definition.")
		logs_putQueryDefinitionCmd.MarkFlagRequired("name")
		logs_putQueryDefinitionCmd.MarkFlagRequired("query-string")
	})
	logsCmd.AddCommand(logs_putQueryDefinitionCmd)
}
