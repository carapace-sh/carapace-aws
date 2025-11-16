package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Gets a resource-based policy of an entity that is identified by its resource ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_getResourcePolicyCmd).Standalone()

	marketplaceCatalog_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the entity resource that is associated with the resource policy.")
	marketplaceCatalog_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_getResourcePolicyCmd)
}
