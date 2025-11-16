package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startMlevaluationTaskRunCmd = &cobra.Command{
	Use:   "start-mlevaluation-task-run",
	Short: "Starts a task to estimate the quality of the transform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startMlevaluationTaskRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_startMlevaluationTaskRunCmd).Standalone()

		glue_startMlevaluationTaskRunCmd.Flags().String("transform-id", "", "The unique identifier of the machine learning transform.")
		glue_startMlevaluationTaskRunCmd.MarkFlagRequired("transform-id")
	})
	glueCmd.AddCommand(glue_startMlevaluationTaskRunCmd)
}
