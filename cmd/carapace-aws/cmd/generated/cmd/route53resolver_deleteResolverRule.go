package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_deleteResolverRuleCmd = &cobra.Command{
	Use:   "delete-resolver-rule",
	Short: "Deletes a Resolver rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_deleteResolverRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_deleteResolverRuleCmd).Standalone()

		route53resolver_deleteResolverRuleCmd.Flags().String("resolver-rule-id", "", "The ID of the Resolver rule that you want to delete.")
		route53resolver_deleteResolverRuleCmd.MarkFlagRequired("resolver-rule-id")
	})
	route53resolverCmd.AddCommand(route53resolver_deleteResolverRuleCmd)
}
