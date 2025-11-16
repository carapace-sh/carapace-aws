package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startExportLabelsTaskRunCmd = &cobra.Command{
	Use:   "start-export-labels-task-run",
	Short: "Begins an asynchronous task to export all labeled data for a particular transform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startExportLabelsTaskRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_startExportLabelsTaskRunCmd).Standalone()

		glue_startExportLabelsTaskRunCmd.Flags().String("output-s3-path", "", "The Amazon S3 path where you export the labels.")
		glue_startExportLabelsTaskRunCmd.Flags().String("transform-id", "", "The unique identifier of the machine learning transform.")
		glue_startExportLabelsTaskRunCmd.MarkFlagRequired("output-s3-path")
		glue_startExportLabelsTaskRunCmd.MarkFlagRequired("transform-id")
	})
	glueCmd.AddCommand(glue_startExportLabelsTaskRunCmd)
}
