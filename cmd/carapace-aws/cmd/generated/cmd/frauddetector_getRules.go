package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getRulesCmd = &cobra.Command{
	Use:   "get-rules",
	Short: "Get all rules for a detector (paginated) if `ruleId` and `ruleVersion` are not specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getRulesCmd).Standalone()

		frauddetector_getRulesCmd.Flags().String("detector-id", "", "The detector ID.")
		frauddetector_getRulesCmd.Flags().String("max-results", "", "The maximum number of rules to return for the request.")
		frauddetector_getRulesCmd.Flags().String("next-token", "", "The next page token.")
		frauddetector_getRulesCmd.Flags().String("rule-id", "", "The rule ID.")
		frauddetector_getRulesCmd.Flags().String("rule-version", "", "The rule version.")
		frauddetector_getRulesCmd.MarkFlagRequired("detector-id")
	})
	frauddetectorCmd.AddCommand(frauddetector_getRulesCmd)
}
