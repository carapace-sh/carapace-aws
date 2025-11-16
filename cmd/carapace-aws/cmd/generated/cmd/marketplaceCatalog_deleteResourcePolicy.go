package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes a resource-based policy on an entity that is identified by its resource ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_deleteResourcePolicyCmd).Standalone()

	marketplaceCatalog_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the entity resource that is associated with the resource policy.")
	marketplaceCatalog_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_deleteResourcePolicyCmd)
}
