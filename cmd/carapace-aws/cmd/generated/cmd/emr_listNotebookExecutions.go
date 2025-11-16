package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listNotebookExecutionsCmd = &cobra.Command{
	Use:   "list-notebook-executions",
	Short: "Provides summaries of all notebook executions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listNotebookExecutionsCmd).Standalone()

	emr_listNotebookExecutionsCmd.Flags().String("editor-id", "", "The unique ID of the editor associated with the notebook execution.")
	emr_listNotebookExecutionsCmd.Flags().String("execution-engine-id", "", "The unique ID of the execution engine.")
	emr_listNotebookExecutionsCmd.Flags().String("from", "", "The beginning of time range filter for listing notebook executions.")
	emr_listNotebookExecutionsCmd.Flags().String("marker", "", "The pagination token, returned by a previous `ListNotebookExecutions` call, that indicates the start of the list for this `ListNotebookExecutions` call.")
	emr_listNotebookExecutionsCmd.Flags().String("status", "", "The status filter for listing notebook executions.")
	emr_listNotebookExecutionsCmd.Flags().String("to", "", "The end of time range filter for listing notebook executions.")
	emrCmd.AddCommand(emr_listNotebookExecutionsCmd)
}
