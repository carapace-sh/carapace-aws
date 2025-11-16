package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listExportJobsCmd = &cobra.Command{
	Use:   "list-export-jobs",
	Short: "Lists all of the export jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listExportJobsCmd).Standalone()

	sesv2_listExportJobsCmd.Flags().String("export-source-type", "", "A value used to list export jobs that have a certain `ExportSourceType`.")
	sesv2_listExportJobsCmd.Flags().String("job-status", "", "A value used to list export jobs that have a certain `JobStatus`.")
	sesv2_listExportJobsCmd.Flags().String("next-token", "", "The pagination token returned from a previous call to `ListExportJobs` to indicate the position in the list of export jobs.")
	sesv2_listExportJobsCmd.Flags().String("page-size", "", "Maximum number of export jobs to return at once.")
	sesv2Cmd.AddCommand(sesv2_listExportJobsCmd)
}
