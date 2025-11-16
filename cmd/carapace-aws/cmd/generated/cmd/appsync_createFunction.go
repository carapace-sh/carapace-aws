package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createFunctionCmd = &cobra.Command{
	Use:   "create-function",
	Short: "Creates a `Function` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createFunctionCmd).Standalone()

	appsync_createFunctionCmd.Flags().String("api-id", "", "The GraphQL API ID.")
	appsync_createFunctionCmd.Flags().String("code", "", "The `function` code that contains the request and response functions.")
	appsync_createFunctionCmd.Flags().String("data-source-name", "", "The `Function` `DataSource` name.")
	appsync_createFunctionCmd.Flags().String("description", "", "The `Function` description.")
	appsync_createFunctionCmd.Flags().String("function-version", "", "The `version` of the request mapping template.")
	appsync_createFunctionCmd.Flags().String("max-batch-size", "", "The maximum batching size for a resolver.")
	appsync_createFunctionCmd.Flags().String("name", "", "The `Function` name.")
	appsync_createFunctionCmd.Flags().String("request-mapping-template", "", "The `Function` request mapping template.")
	appsync_createFunctionCmd.Flags().String("response-mapping-template", "", "The `Function` response mapping template.")
	appsync_createFunctionCmd.Flags().String("runtime", "", "")
	appsync_createFunctionCmd.Flags().String("sync-config", "", "")
	appsync_createFunctionCmd.MarkFlagRequired("api-id")
	appsync_createFunctionCmd.MarkFlagRequired("data-source-name")
	appsync_createFunctionCmd.MarkFlagRequired("name")
	appsyncCmd.AddCommand(appsync_createFunctionCmd)
}
