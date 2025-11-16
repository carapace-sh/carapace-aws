package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteFunctionCmd = &cobra.Command{
	Use:   "delete-function",
	Short: "Deletes a `Function`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteFunctionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_deleteFunctionCmd).Standalone()

		appsync_deleteFunctionCmd.Flags().String("api-id", "", "The GraphQL API ID.")
		appsync_deleteFunctionCmd.Flags().String("function-id", "", "The `Function` ID.")
		appsync_deleteFunctionCmd.MarkFlagRequired("api-id")
		appsync_deleteFunctionCmd.MarkFlagRequired("function-id")
	})
	appsyncCmd.AddCommand(appsync_deleteFunctionCmd)
}
