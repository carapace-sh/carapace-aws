package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_startConfigRulesEvaluationCmd = &cobra.Command{
	Use:   "start-config-rules-evaluation",
	Short: "Runs an on-demand evaluation for the specified Config rules against the last known configuration state of the resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_startConfigRulesEvaluationCmd).Standalone()

	config_startConfigRulesEvaluationCmd.Flags().String("config-rule-names", "", "The list of names of Config rules that you want to run evaluations for.")
	configCmd.AddCommand(config_startConfigRulesEvaluationCmd)
}
