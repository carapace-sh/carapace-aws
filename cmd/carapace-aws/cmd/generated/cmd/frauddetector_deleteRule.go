package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_deleteRuleCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "Deletes the rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_deleteRuleCmd).Standalone()

	frauddetector_deleteRuleCmd.Flags().String("rule", "", "")
	frauddetector_deleteRuleCmd.MarkFlagRequired("rule")
	frauddetectorCmd.AddCommand(frauddetector_deleteRuleCmd)
}
