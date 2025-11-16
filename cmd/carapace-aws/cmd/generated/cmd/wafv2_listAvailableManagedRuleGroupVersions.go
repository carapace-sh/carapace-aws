package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_listAvailableManagedRuleGroupVersionsCmd = &cobra.Command{
	Use:   "list-available-managed-rule-group-versions",
	Short: "Returns a list of the available versions for the specified managed rule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_listAvailableManagedRuleGroupVersionsCmd).Standalone()

	wafv2_listAvailableManagedRuleGroupVersionsCmd.Flags().String("limit", "", "The maximum number of objects that you want WAF to return for this request.")
	wafv2_listAvailableManagedRuleGroupVersionsCmd.Flags().String("name", "", "The name of the managed rule group.")
	wafv2_listAvailableManagedRuleGroupVersionsCmd.Flags().String("next-marker", "", "When you request a list of objects with a `Limit` setting, if the number of objects that are still available for retrieval exceeds the limit, WAF returns a `NextMarker` value in the response.")
	wafv2_listAvailableManagedRuleGroupVersionsCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_listAvailableManagedRuleGroupVersionsCmd.Flags().String("vendor-name", "", "The name of the managed rule group vendor.")
	wafv2_listAvailableManagedRuleGroupVersionsCmd.MarkFlagRequired("name")
	wafv2_listAvailableManagedRuleGroupVersionsCmd.MarkFlagRequired("scope")
	wafv2_listAvailableManagedRuleGroupVersionsCmd.MarkFlagRequired("vendor-name")
	wafv2Cmd.AddCommand(wafv2_listAvailableManagedRuleGroupVersionsCmd)
}
