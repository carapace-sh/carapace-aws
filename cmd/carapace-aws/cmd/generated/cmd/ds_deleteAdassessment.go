package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_deleteAdassessmentCmd = &cobra.Command{
	Use:   "delete-adassessment",
	Short: "Deletes a directory assessment and all associated data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_deleteAdassessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_deleteAdassessmentCmd).Standalone()

		ds_deleteAdassessmentCmd.Flags().String("assessment-id", "", "The unique identifier of the directory assessment to delete.")
		ds_deleteAdassessmentCmd.MarkFlagRequired("assessment-id")
	})
	dsCmd.AddCommand(ds_deleteAdassessmentCmd)
}
