package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_updateRuleMetadataCmd = &cobra.Command{
	Use:   "update-rule-metadata",
	Short: "Updates a rule's metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_updateRuleMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_updateRuleMetadataCmd).Standalone()

		frauddetector_updateRuleMetadataCmd.Flags().String("description", "", "The rule description.")
		frauddetector_updateRuleMetadataCmd.Flags().String("rule", "", "The rule to update.")
		frauddetector_updateRuleMetadataCmd.MarkFlagRequired("description")
		frauddetector_updateRuleMetadataCmd.MarkFlagRequired("rule")
	})
	frauddetectorCmd.AddCommand(frauddetector_updateRuleMetadataCmd)
}
