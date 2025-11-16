package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_putGraphqlApiEnvironmentVariablesCmd = &cobra.Command{
	Use:   "put-graphql-api-environment-variables",
	Short: "Creates a list of environmental variables in an API by its ID value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_putGraphqlApiEnvironmentVariablesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_putGraphqlApiEnvironmentVariablesCmd).Standalone()

		appsync_putGraphqlApiEnvironmentVariablesCmd.Flags().String("api-id", "", "The ID of the API to which the environmental variable list will be written.")
		appsync_putGraphqlApiEnvironmentVariablesCmd.Flags().String("environment-variables", "", "The list of environmental variables to add to the API.")
		appsync_putGraphqlApiEnvironmentVariablesCmd.MarkFlagRequired("api-id")
		appsync_putGraphqlApiEnvironmentVariablesCmd.MarkFlagRequired("environment-variables")
	})
	appsyncCmd.AddCommand(appsync_putGraphqlApiEnvironmentVariablesCmd)
}
