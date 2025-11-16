package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getGraphqlApiEnvironmentVariablesCmd = &cobra.Command{
	Use:   "get-graphql-api-environment-variables",
	Short: "Retrieves the list of environmental variable key-value pairs associated with an API by its ID value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getGraphqlApiEnvironmentVariablesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getGraphqlApiEnvironmentVariablesCmd).Standalone()

		appsync_getGraphqlApiEnvironmentVariablesCmd.Flags().String("api-id", "", "The ID of the API from which the environmental variable list will be retrieved.")
		appsync_getGraphqlApiEnvironmentVariablesCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_getGraphqlApiEnvironmentVariablesCmd)
}
