package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceCatalog_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based policy to an entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceCatalog_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceCatalog_putResourcePolicyCmd).Standalone()

		marketplaceCatalog_putResourcePolicyCmd.Flags().String("policy", "", "The policy document to set; formatted in JSON.")
		marketplaceCatalog_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the entity resource you want to associate with a resource policy.")
		marketplaceCatalog_putResourcePolicyCmd.MarkFlagRequired("policy")
		marketplaceCatalog_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	marketplaceCatalogCmd.AddCommand(marketplaceCatalog_putResourcePolicyCmd)
}
