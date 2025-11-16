package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listNodeFromTemplateJobsCmd = &cobra.Command{
	Use:   "list-node-from-template-jobs",
	Short: "Returns a list of camera stream node jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listNodeFromTemplateJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_listNodeFromTemplateJobsCmd).Standalone()

		panorama_listNodeFromTemplateJobsCmd.Flags().String("max-results", "", "The maximum number of node from template jobs to return in one page of results.")
		panorama_listNodeFromTemplateJobsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	})
	panoramaCmd.AddCommand(panorama_listNodeFromTemplateJobsCmd)
}
