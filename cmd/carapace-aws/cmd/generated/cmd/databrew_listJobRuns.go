package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_listJobRunsCmd = &cobra.Command{
	Use:   "list-job-runs",
	Short: "Lists all of the previous runs of a particular DataBrew job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_listJobRunsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_listJobRunsCmd).Standalone()

		databrew_listJobRunsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
		databrew_listJobRunsCmd.Flags().String("name", "", "The name of the job.")
		databrew_listJobRunsCmd.Flags().String("next-token", "", "The token returned by a previous call to retrieve the next set of results.")
		databrew_listJobRunsCmd.MarkFlagRequired("name")
	})
	databrewCmd.AddCommand(databrew_listJobRunsCmd)
}
