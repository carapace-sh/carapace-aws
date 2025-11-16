package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_getJobsQueryResultsCmd = &cobra.Command{
	Use:   "get-jobs-query-results",
	Short: "Retrieve a JSON array of up to twenty of your most recent jobs matched by a jobs query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_getJobsQueryResultsCmd).Standalone()

	mediaconvert_getJobsQueryResultsCmd.Flags().String("id", "", "The ID of the jobs query.")
	mediaconvert_getJobsQueryResultsCmd.MarkFlagRequired("id")
	mediaconvertCmd.AddCommand(mediaconvert_getJobsQueryResultsCmd)
}
