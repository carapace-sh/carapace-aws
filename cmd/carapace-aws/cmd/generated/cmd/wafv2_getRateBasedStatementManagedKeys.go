package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getRateBasedStatementManagedKeysCmd = &cobra.Command{
	Use:   "get-rate-based-statement-managed-keys",
	Short: "Retrieves the IP addresses that are currently blocked by a rate-based rule instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getRateBasedStatementManagedKeysCmd).Standalone()

	wafv2_getRateBasedStatementManagedKeysCmd.Flags().String("rule-group-rule-name", "", "The name of the rule group reference statement in your web ACL.")
	wafv2_getRateBasedStatementManagedKeysCmd.Flags().String("rule-name", "", "The name of the rate-based rule to get the keys for.")
	wafv2_getRateBasedStatementManagedKeysCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_getRateBasedStatementManagedKeysCmd.Flags().String("web-aclid", "", "The unique identifier for the web ACL.")
	wafv2_getRateBasedStatementManagedKeysCmd.Flags().String("web-aclname", "", "The name of the web ACL.")
	wafv2_getRateBasedStatementManagedKeysCmd.MarkFlagRequired("rule-name")
	wafv2_getRateBasedStatementManagedKeysCmd.MarkFlagRequired("scope")
	wafv2_getRateBasedStatementManagedKeysCmd.MarkFlagRequired("web-aclid")
	wafv2_getRateBasedStatementManagedKeysCmd.MarkFlagRequired("web-aclname")
	wafv2Cmd.AddCommand(wafv2_getRateBasedStatementManagedKeysCmd)
}
