package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listResolverRuleAssociationsCmd = &cobra.Command{
	Use:   "list-resolver-rule-associations",
	Short: "Lists the associations that were created between Resolver rules and VPCs using the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listResolverRuleAssociationsCmd).Standalone()

	route53resolver_listResolverRuleAssociationsCmd.Flags().String("filters", "", "An optional specification to return a subset of Resolver rules, such as Resolver rules that are associated with the same VPC ID.")
	route53resolver_listResolverRuleAssociationsCmd.Flags().String("max-results", "", "The maximum number of rule associations that you want to return in the response to a `ListResolverRuleAssociations` request.")
	route53resolver_listResolverRuleAssociationsCmd.Flags().String("next-token", "", "For the first `ListResolverRuleAssociation` request, omit this value.")
	route53resolverCmd.AddCommand(route53resolver_listResolverRuleAssociationsCmd)
}
