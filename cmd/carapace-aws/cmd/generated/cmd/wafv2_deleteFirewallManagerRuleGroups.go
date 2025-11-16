package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_deleteFirewallManagerRuleGroupsCmd = &cobra.Command{
	Use:   "delete-firewall-manager-rule-groups",
	Short: "Deletes all rule groups that are managed by Firewall Manager from the specified [WebACL]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_deleteFirewallManagerRuleGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_deleteFirewallManagerRuleGroupsCmd).Standalone()

		wafv2_deleteFirewallManagerRuleGroupsCmd.Flags().String("web-aclarn", "", "The Amazon Resource Name (ARN) of the web ACL.")
		wafv2_deleteFirewallManagerRuleGroupsCmd.Flags().String("web-acllock-token", "", "A token used for optimistic locking.")
		wafv2_deleteFirewallManagerRuleGroupsCmd.MarkFlagRequired("web-aclarn")
		wafv2_deleteFirewallManagerRuleGroupsCmd.MarkFlagRequired("web-acllock-token")
	})
	wafv2Cmd.AddCommand(wafv2_deleteFirewallManagerRuleGroupsCmd)
}
