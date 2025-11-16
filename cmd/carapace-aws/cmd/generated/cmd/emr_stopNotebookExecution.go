package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_stopNotebookExecutionCmd = &cobra.Command{
	Use:   "stop-notebook-execution",
	Short: "Stops a notebook execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_stopNotebookExecutionCmd).Standalone()

	emr_stopNotebookExecutionCmd.Flags().String("notebook-execution-id", "", "The unique identifier of the notebook execution.")
	emr_stopNotebookExecutionCmd.MarkFlagRequired("notebook-execution-id")
	emrCmd.AddCommand(emr_stopNotebookExecutionCmd)
}
