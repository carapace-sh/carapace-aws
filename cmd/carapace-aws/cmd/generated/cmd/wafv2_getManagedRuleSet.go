package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_getManagedRuleSetCmd = &cobra.Command{
	Use:   "get-managed-rule-set",
	Short: "Retrieves the specified managed rule set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_getManagedRuleSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_getManagedRuleSetCmd).Standalone()

		wafv2_getManagedRuleSetCmd.Flags().String("id", "", "A unique identifier for the managed rule set.")
		wafv2_getManagedRuleSetCmd.Flags().String("name", "", "The name of the managed rule set.")
		wafv2_getManagedRuleSetCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_getManagedRuleSetCmd.MarkFlagRequired("id")
		wafv2_getManagedRuleSetCmd.MarkFlagRequired("name")
		wafv2_getManagedRuleSetCmd.MarkFlagRequired("scope")
	})
	wafv2Cmd.AddCommand(wafv2_getManagedRuleSetCmd)
}
