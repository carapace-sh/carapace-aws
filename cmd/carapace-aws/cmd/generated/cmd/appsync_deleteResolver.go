package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_deleteResolverCmd = &cobra.Command{
	Use:   "delete-resolver",
	Short: "Deletes a `Resolver` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_deleteResolverCmd).Standalone()

	appsync_deleteResolverCmd.Flags().String("api-id", "", "The API ID.")
	appsync_deleteResolverCmd.Flags().String("field-name", "", "The resolver field name.")
	appsync_deleteResolverCmd.Flags().String("type-name", "", "The name of the resolver type.")
	appsync_deleteResolverCmd.MarkFlagRequired("api-id")
	appsync_deleteResolverCmd.MarkFlagRequired("field-name")
	appsync_deleteResolverCmd.MarkFlagRequired("type-name")
	appsyncCmd.AddCommand(appsync_deleteResolverCmd)
}
