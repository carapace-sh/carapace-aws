package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_cancelMltaskRunCmd = &cobra.Command{
	Use:   "cancel-mltask-run",
	Short: "Cancels (stops) a task run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_cancelMltaskRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_cancelMltaskRunCmd).Standalone()

		glue_cancelMltaskRunCmd.Flags().String("task-run-id", "", "A unique identifier for the task run.")
		glue_cancelMltaskRunCmd.Flags().String("transform-id", "", "The unique identifier of the machine learning transform.")
		glue_cancelMltaskRunCmd.MarkFlagRequired("task-run-id")
		glue_cancelMltaskRunCmd.MarkFlagRequired("transform-id")
	})
	glueCmd.AddCommand(glue_cancelMltaskRunCmd)
}
