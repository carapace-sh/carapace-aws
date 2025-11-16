package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getSensitivityInspectionTemplateCmd = &cobra.Command{
	Use:   "get-sensitivity-inspection-template",
	Short: "Retrieves the settings for the sensitivity inspection template for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getSensitivityInspectionTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getSensitivityInspectionTemplateCmd).Standalone()

		macie2_getSensitivityInspectionTemplateCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
		macie2_getSensitivityInspectionTemplateCmd.MarkFlagRequired("id")
	})
	macie2Cmd.AddCommand(macie2_getSensitivityInspectionTemplateCmd)
}
