package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getResolverCmd = &cobra.Command{
	Use:   "get-resolver",
	Short: "Retrieves a `Resolver` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getResolverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getResolverCmd).Standalone()

		appsync_getResolverCmd.Flags().String("api-id", "", "The API ID.")
		appsync_getResolverCmd.Flags().String("field-name", "", "The resolver field name.")
		appsync_getResolverCmd.Flags().String("type-name", "", "The resolver type name.")
		appsync_getResolverCmd.MarkFlagRequired("api-id")
		appsync_getResolverCmd.MarkFlagRequired("field-name")
		appsync_getResolverCmd.MarkFlagRequired("type-name")
	})
	appsyncCmd.AddCommand(appsync_getResolverCmd)
}
