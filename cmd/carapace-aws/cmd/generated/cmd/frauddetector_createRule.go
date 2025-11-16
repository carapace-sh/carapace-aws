package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_createRuleCmd = &cobra.Command{
	Use:   "create-rule",
	Short: "Creates a rule for use with the specified detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_createRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_createRuleCmd).Standalone()

		frauddetector_createRuleCmd.Flags().String("description", "", "The rule description.")
		frauddetector_createRuleCmd.Flags().String("detector-id", "", "The detector ID for the rule's parent detector.")
		frauddetector_createRuleCmd.Flags().String("expression", "", "The rule expression.")
		frauddetector_createRuleCmd.Flags().String("language", "", "The language of the rule.")
		frauddetector_createRuleCmd.Flags().String("outcomes", "", "The outcome or outcomes returned when the rule expression matches.")
		frauddetector_createRuleCmd.Flags().String("rule-id", "", "The rule ID.")
		frauddetector_createRuleCmd.Flags().String("tags", "", "A collection of key and value pairs.")
		frauddetector_createRuleCmd.MarkFlagRequired("detector-id")
		frauddetector_createRuleCmd.MarkFlagRequired("expression")
		frauddetector_createRuleCmd.MarkFlagRequired("language")
		frauddetector_createRuleCmd.MarkFlagRequired("outcomes")
		frauddetector_createRuleCmd.MarkFlagRequired("rule-id")
	})
	frauddetectorCmd.AddCommand(frauddetector_createRuleCmd)
}
