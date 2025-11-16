package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putConfigRuleCmd = &cobra.Command{
	Use:   "put-config-rule",
	Short: "Adds or updates an Config rule to evaluate if your Amazon Web Services resources comply with your desired configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putConfigRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_putConfigRuleCmd).Standalone()

		config_putConfigRuleCmd.Flags().String("config-rule", "", "The rule that you want to add to your account.")
		config_putConfigRuleCmd.Flags().String("tags", "", "An array of tag object.")
		config_putConfigRuleCmd.MarkFlagRequired("config-rule")
	})
	configCmd.AddCommand(config_putConfigRuleCmd)
}
