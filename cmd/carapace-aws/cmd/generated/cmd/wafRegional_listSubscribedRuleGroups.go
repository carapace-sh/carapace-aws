package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listSubscribedRuleGroupsCmd = &cobra.Command{
	Use:   "list-subscribed-rule-groups",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listSubscribedRuleGroupsCmd).Standalone()

	wafRegional_listSubscribedRuleGroupsCmd.Flags().String("limit", "", "Specifies the number of subscribed rule groups that you want AWS WAF to return for this request.")
	wafRegional_listSubscribedRuleGroupsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `ByteMatchSets`subscribed rule groups than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of subscribed rule groups.")
	wafRegionalCmd.AddCommand(wafRegional_listSubscribedRuleGroupsCmd)
}
