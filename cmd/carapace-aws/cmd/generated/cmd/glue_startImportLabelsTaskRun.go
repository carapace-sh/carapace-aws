package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startImportLabelsTaskRunCmd = &cobra.Command{
	Use:   "start-import-labels-task-run",
	Short: "Enables you to provide additional labels (examples of truth) to be used to teach the machine learning transform and improve its quality.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startImportLabelsTaskRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_startImportLabelsTaskRunCmd).Standalone()

		glue_startImportLabelsTaskRunCmd.Flags().String("input-s3-path", "", "The Amazon Simple Storage Service (Amazon S3) path from where you import the labels.")
		glue_startImportLabelsTaskRunCmd.Flags().String("replace-all-labels", "", "Indicates whether to overwrite your existing labels.")
		glue_startImportLabelsTaskRunCmd.Flags().String("transform-id", "", "The unique identifier of the machine learning transform.")
		glue_startImportLabelsTaskRunCmd.MarkFlagRequired("input-s3-path")
		glue_startImportLabelsTaskRunCmd.MarkFlagRequired("transform-id")
	})
	glueCmd.AddCommand(glue_startImportLabelsTaskRunCmd)
}
