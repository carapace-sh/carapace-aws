package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateFunctionCmd = &cobra.Command{
	Use:   "update-function",
	Short: "Updates a `Function` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateFunctionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_updateFunctionCmd).Standalone()

		appsync_updateFunctionCmd.Flags().String("api-id", "", "The GraphQL API ID.")
		appsync_updateFunctionCmd.Flags().String("code", "", "The `function` code that contains the request and response functions.")
		appsync_updateFunctionCmd.Flags().String("data-source-name", "", "The `Function` `DataSource` name.")
		appsync_updateFunctionCmd.Flags().String("description", "", "The `Function` description.")
		appsync_updateFunctionCmd.Flags().String("function-id", "", "The function ID.")
		appsync_updateFunctionCmd.Flags().String("function-version", "", "The `version` of the request mapping template.")
		appsync_updateFunctionCmd.Flags().String("max-batch-size", "", "The maximum batching size for a resolver.")
		appsync_updateFunctionCmd.Flags().String("name", "", "The `Function` name.")
		appsync_updateFunctionCmd.Flags().String("request-mapping-template", "", "The `Function` request mapping template.")
		appsync_updateFunctionCmd.Flags().String("response-mapping-template", "", "The `Function` request mapping template.")
		appsync_updateFunctionCmd.Flags().String("runtime", "", "")
		appsync_updateFunctionCmd.Flags().String("sync-config", "", "")
		appsync_updateFunctionCmd.MarkFlagRequired("api-id")
		appsync_updateFunctionCmd.MarkFlagRequired("data-source-name")
		appsync_updateFunctionCmd.MarkFlagRequired("function-id")
		appsync_updateFunctionCmd.MarkFlagRequired("name")
	})
	appsyncCmd.AddCommand(appsync_updateFunctionCmd)
}
