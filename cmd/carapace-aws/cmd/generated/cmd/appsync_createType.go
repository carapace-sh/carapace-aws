package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createTypeCmd = &cobra.Command{
	Use:   "create-type",
	Short: "Creates a `Type` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_createTypeCmd).Standalone()

		appsync_createTypeCmd.Flags().String("api-id", "", "The API ID.")
		appsync_createTypeCmd.Flags().String("definition", "", "The type definition, in GraphQL Schema Definition Language (SDL) format.")
		appsync_createTypeCmd.Flags().String("format", "", "The type format: SDL or JSON.")
		appsync_createTypeCmd.MarkFlagRequired("api-id")
		appsync_createTypeCmd.MarkFlagRequired("definition")
		appsync_createTypeCmd.MarkFlagRequired("format")
	})
	appsyncCmd.AddCommand(appsync_createTypeCmd)
}
