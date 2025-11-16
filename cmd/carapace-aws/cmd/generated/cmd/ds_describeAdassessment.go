package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeAdassessmentCmd = &cobra.Command{
	Use:   "describe-adassessment",
	Short: "Retrieves detailed information about a directory assessment, including its current status, validation results, and configuration details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeAdassessmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_describeAdassessmentCmd).Standalone()

		ds_describeAdassessmentCmd.Flags().String("assessment-id", "", "The identifier of the directory assessment to describe.")
		ds_describeAdassessmentCmd.MarkFlagRequired("assessment-id")
	})
	dsCmd.AddCommand(ds_describeAdassessmentCmd)
}
