package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_associateSourceGraphqlApiCmd = &cobra.Command{
	Use:   "associate-source-graphql-api",
	Short: "Creates an association between a Merged API and source API using the Merged API's identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_associateSourceGraphqlApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_associateSourceGraphqlApiCmd).Standalone()

		appsync_associateSourceGraphqlApiCmd.Flags().String("description", "", "The description field.")
		appsync_associateSourceGraphqlApiCmd.Flags().String("merged-api-identifier", "", "The identifier of the AppSync Merged API.")
		appsync_associateSourceGraphqlApiCmd.Flags().String("source-api-association-config", "", "The `SourceApiAssociationConfig` object data.")
		appsync_associateSourceGraphqlApiCmd.Flags().String("source-api-identifier", "", "The identifier of the AppSync Source API.")
		appsync_associateSourceGraphqlApiCmd.MarkFlagRequired("merged-api-identifier")
		appsync_associateSourceGraphqlApiCmd.MarkFlagRequired("source-api-identifier")
	})
	appsyncCmd.AddCommand(appsync_associateSourceGraphqlApiCmd)
}
