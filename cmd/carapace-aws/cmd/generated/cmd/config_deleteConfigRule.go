package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteConfigRuleCmd = &cobra.Command{
	Use:   "delete-config-rule",
	Short: "Deletes the specified Config rule and all of its evaluation results.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteConfigRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deleteConfigRuleCmd).Standalone()

		config_deleteConfigRuleCmd.Flags().String("config-rule-name", "", "The name of the Config rule that you want to delete.")
		config_deleteConfigRuleCmd.MarkFlagRequired("config-rule-name")
	})
	configCmd.AddCommand(config_deleteConfigRuleCmd)
}
