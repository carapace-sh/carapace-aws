package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_getTargetSelectionRulesCmd = &cobra.Command{
	Use:   "get-target-selection-rules",
	Short: "Converts source selection rules into their target counterparts for schema conversion operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_getTargetSelectionRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_getTargetSelectionRulesCmd).Standalone()

		dms_getTargetSelectionRulesCmd.Flags().String("migration-project-identifier", "", "The migration project name or Amazon Resource Name (ARN).")
		dms_getTargetSelectionRulesCmd.Flags().String("selection-rules", "", "The JSON string representing the source selection rules for conversion.")
		dms_getTargetSelectionRulesCmd.MarkFlagRequired("migration-project-identifier")
		dms_getTargetSelectionRulesCmd.MarkFlagRequired("selection-rules")
	})
	dmsCmd.AddCommand(dms_getTargetSelectionRulesCmd)
}
