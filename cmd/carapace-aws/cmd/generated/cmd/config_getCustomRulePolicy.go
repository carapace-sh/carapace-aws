package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getCustomRulePolicyCmd = &cobra.Command{
	Use:   "get-custom-rule-policy",
	Short: "Returns the policy definition containing the logic for your Config Custom Policy rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCustomRulePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getCustomRulePolicyCmd).Standalone()

		config_getCustomRulePolicyCmd.Flags().String("config-rule-name", "", "The name of your Config Custom Policy rule.")
	})
	configCmd.AddCommand(config_getCustomRulePolicyCmd)
}
