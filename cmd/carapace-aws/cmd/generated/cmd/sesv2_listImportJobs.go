package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listImportJobsCmd = &cobra.Command{
	Use:   "list-import-jobs",
	Short: "Lists all of the import jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listImportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_listImportJobsCmd).Standalone()

		sesv2_listImportJobsCmd.Flags().String("import-destination-type", "", "The destination of the import job, which can be used to list import jobs that have a certain `ImportDestinationType`.")
		sesv2_listImportJobsCmd.Flags().String("next-token", "", "A string token indicating that there might be additional import jobs available to be listed.")
		sesv2_listImportJobsCmd.Flags().String("page-size", "", "Maximum number of import jobs to return at once.")
	})
	sesv2Cmd.AddCommand(sesv2_listImportJobsCmd)
}
