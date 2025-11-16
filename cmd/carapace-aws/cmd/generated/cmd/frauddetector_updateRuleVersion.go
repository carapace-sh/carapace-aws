package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateRuleVersionCmd = &cobra.Command{
	Use:   "update-rule-version",
	Short: "Updates a rule version resulting in a new rule version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateRuleVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_updateRuleVersionCmd).Standalone()

		frauddetector_updateRuleVersionCmd.Flags().String("description", "", "The description.")
		frauddetector_updateRuleVersionCmd.Flags().String("expression", "", "The rule expression.")
		frauddetector_updateRuleVersionCmd.Flags().String("language", "", "The language.")
		frauddetector_updateRuleVersionCmd.Flags().String("outcomes", "", "The outcomes.")
		frauddetector_updateRuleVersionCmd.Flags().String("rule", "", "The rule to update.")
		frauddetector_updateRuleVersionCmd.Flags().String("tags", "", "The tags to assign to the rule version.")
		frauddetector_updateRuleVersionCmd.MarkFlagRequired("expression")
		frauddetector_updateRuleVersionCmd.MarkFlagRequired("language")
		frauddetector_updateRuleVersionCmd.MarkFlagRequired("outcomes")
		frauddetector_updateRuleVersionCmd.MarkFlagRequired("rule")
	})
	frauddetectorCmd.AddCommand(frauddetector_updateRuleVersionCmd)
}
