package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getFunctionCmd = &cobra.Command{
	Use:   "get-function",
	Short: "Get a `Function`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getFunctionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getFunctionCmd).Standalone()

		appsync_getFunctionCmd.Flags().String("api-id", "", "The GraphQL API ID.")
		appsync_getFunctionCmd.Flags().String("function-id", "", "The `Function` ID.")
		appsync_getFunctionCmd.MarkFlagRequired("api-id")
		appsync_getFunctionCmd.MarkFlagRequired("function-id")
	})
	appsyncCmd.AddCommand(appsync_getFunctionCmd)
}
