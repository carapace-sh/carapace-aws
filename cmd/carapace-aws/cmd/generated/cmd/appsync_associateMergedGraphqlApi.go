package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_associateMergedGraphqlApiCmd = &cobra.Command{
	Use:   "associate-merged-graphql-api",
	Short: "Creates an association between a Merged API and source API using the source API's identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_associateMergedGraphqlApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_associateMergedGraphqlApiCmd).Standalone()

		appsync_associateMergedGraphqlApiCmd.Flags().String("description", "", "The description field.")
		appsync_associateMergedGraphqlApiCmd.Flags().String("merged-api-identifier", "", "The identifier of the AppSync Merged API.")
		appsync_associateMergedGraphqlApiCmd.Flags().String("source-api-association-config", "", "The `SourceApiAssociationConfig` object data.")
		appsync_associateMergedGraphqlApiCmd.Flags().String("source-api-identifier", "", "The identifier of the AppSync Source API.")
		appsync_associateMergedGraphqlApiCmd.MarkFlagRequired("merged-api-identifier")
		appsync_associateMergedGraphqlApiCmd.MarkFlagRequired("source-api-identifier")
	})
	appsyncCmd.AddCommand(appsync_associateMergedGraphqlApiCmd)
}
