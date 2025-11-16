package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_putManagedRuleSetVersionsCmd = &cobra.Command{
	Use:   "put-managed-rule-set-versions",
	Short: "Defines the versions of your managed rule set that you are offering to the customers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_putManagedRuleSetVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_putManagedRuleSetVersionsCmd).Standalone()

		wafv2_putManagedRuleSetVersionsCmd.Flags().String("id", "", "A unique identifier for the managed rule set.")
		wafv2_putManagedRuleSetVersionsCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
		wafv2_putManagedRuleSetVersionsCmd.Flags().String("name", "", "The name of the managed rule set.")
		wafv2_putManagedRuleSetVersionsCmd.Flags().String("recommended-version", "", "The version of the named managed rule group that you'd like your customers to choose, from among your version offerings.")
		wafv2_putManagedRuleSetVersionsCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_putManagedRuleSetVersionsCmd.Flags().String("versions-to-publish", "", "The versions of the named managed rule group that you want to offer to your customers.")
		wafv2_putManagedRuleSetVersionsCmd.MarkFlagRequired("id")
		wafv2_putManagedRuleSetVersionsCmd.MarkFlagRequired("lock-token")
		wafv2_putManagedRuleSetVersionsCmd.MarkFlagRequired("name")
		wafv2_putManagedRuleSetVersionsCmd.MarkFlagRequired("scope")
	})
	wafv2Cmd.AddCommand(wafv2_putManagedRuleSetVersionsCmd)
}
