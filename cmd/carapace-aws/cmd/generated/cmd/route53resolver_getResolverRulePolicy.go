package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getResolverRulePolicyCmd = &cobra.Command{
	Use:   "get-resolver-rule-policy",
	Short: "Gets information about the Resolver rule policy for a specified rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getResolverRulePolicyCmd).Standalone()

	route53resolver_getResolverRulePolicyCmd.Flags().String("arn", "", "The ID of the Resolver rule that you want to get the Resolver rule policy for.")
	route53resolver_getResolverRulePolicyCmd.MarkFlagRequired("arn")
	route53resolverCmd.AddCommand(route53resolver_getResolverRulePolicyCmd)
}
