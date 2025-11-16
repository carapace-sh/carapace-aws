package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateClassificationJobCmd = &cobra.Command{
	Use:   "update-classification-job",
	Short: "Changes the status of a classification job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateClassificationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_updateClassificationJobCmd).Standalone()

		macie2_updateClassificationJobCmd.Flags().String("job-id", "", "The unique identifier for the classification job.")
		macie2_updateClassificationJobCmd.Flags().String("job-status", "", "The new status for the job.")
		macie2_updateClassificationJobCmd.MarkFlagRequired("job-id")
		macie2_updateClassificationJobCmd.MarkFlagRequired("job-status")
	})
	macie2Cmd.AddCommand(macie2_updateClassificationJobCmd)
}
