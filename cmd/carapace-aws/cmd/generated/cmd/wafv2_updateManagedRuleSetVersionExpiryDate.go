package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_updateManagedRuleSetVersionExpiryDateCmd = &cobra.Command{
	Use:   "update-managed-rule-set-version-expiry-date",
	Short: "Updates the expiration information for your managed rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_updateManagedRuleSetVersionExpiryDateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_updateManagedRuleSetVersionExpiryDateCmd).Standalone()

		wafv2_updateManagedRuleSetVersionExpiryDateCmd.Flags().String("expiry-timestamp", "", "The time that you want the version to expire.")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.Flags().String("id", "", "A unique identifier for the managed rule set.")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.Flags().String("name", "", "The name of the managed rule set.")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.Flags().String("version-to-expire", "", "The version that you want to remove from your list of offerings for the named managed rule group.")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.MarkFlagRequired("expiry-timestamp")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.MarkFlagRequired("id")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.MarkFlagRequired("lock-token")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.MarkFlagRequired("name")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.MarkFlagRequired("scope")
		wafv2_updateManagedRuleSetVersionExpiryDateCmd.MarkFlagRequired("version-to-expire")
	})
	wafv2Cmd.AddCommand(wafv2_updateManagedRuleSetVersionExpiryDateCmd)
}
