package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateSensitivityInspectionTemplateCmd = &cobra.Command{
	Use:   "update-sensitivity-inspection-template",
	Short: "Updates the settings for the sensitivity inspection template for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateSensitivityInspectionTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_updateSensitivityInspectionTemplateCmd).Standalone()

		macie2_updateSensitivityInspectionTemplateCmd.Flags().String("description", "", "A custom description of the template.")
		macie2_updateSensitivityInspectionTemplateCmd.Flags().String("excludes", "", "The managed data identifiers to explicitly exclude (not use) when performing automated sensitive data discovery.")
		macie2_updateSensitivityInspectionTemplateCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
		macie2_updateSensitivityInspectionTemplateCmd.Flags().String("includes", "", "The allow lists, custom data identifiers, and managed data identifiers to explicitly include (use) when performing automated sensitive data discovery.")
		macie2_updateSensitivityInspectionTemplateCmd.MarkFlagRequired("id")
	})
	macie2Cmd.AddCommand(macie2_updateSensitivityInspectionTemplateCmd)
}
