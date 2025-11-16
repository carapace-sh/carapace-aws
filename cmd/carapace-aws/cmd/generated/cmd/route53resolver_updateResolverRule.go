package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_updateResolverRuleCmd = &cobra.Command{
	Use:   "update-resolver-rule",
	Short: "Updates settings for a specified Resolver rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_updateResolverRuleCmd).Standalone()

	route53resolver_updateResolverRuleCmd.Flags().String("config", "", "The new settings for the Resolver rule.")
	route53resolver_updateResolverRuleCmd.Flags().String("resolver-rule-id", "", "The ID of the Resolver rule that you want to update.")
	route53resolver_updateResolverRuleCmd.MarkFlagRequired("config")
	route53resolver_updateResolverRuleCmd.MarkFlagRequired("resolver-rule-id")
	route53resolverCmd.AddCommand(route53resolver_updateResolverRuleCmd)
}
