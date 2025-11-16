package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_describeNotebookExecutionCmd = &cobra.Command{
	Use:   "describe-notebook-execution",
	Short: "Provides details of a notebook execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_describeNotebookExecutionCmd).Standalone()

	emr_describeNotebookExecutionCmd.Flags().String("notebook-execution-id", "", "The unique identifier of the notebook execution.")
	emr_describeNotebookExecutionCmd.MarkFlagRequired("notebook-execution-id")
	emrCmd.AddCommand(emr_describeNotebookExecutionCmd)
}
