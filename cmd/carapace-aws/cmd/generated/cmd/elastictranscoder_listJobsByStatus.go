package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_listJobsByStatusCmd = &cobra.Command{
	Use:   "list-jobs-by-status",
	Short: "The ListJobsByStatus operation gets a list of jobs that have a specified status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_listJobsByStatusCmd).Standalone()

	elastictranscoder_listJobsByStatusCmd.Flags().String("ascending", "", "To list jobs in chronological order by the date and time that they were submitted, enter `true`.")
	elastictranscoder_listJobsByStatusCmd.Flags().String("page-token", "", "When Elastic Transcoder returns more than one page of results, use `pageToken` in subsequent `GET` requests to get each successive page of results.")
	elastictranscoder_listJobsByStatusCmd.Flags().String("status", "", "To get information about all of the jobs associated with the current AWS account that have a given status, specify the following status: `Submitted`, `Progressing`, `Complete`, `Canceled`, or `Error`.")
	elastictranscoder_listJobsByStatusCmd.MarkFlagRequired("status")
	elastictranscoderCmd.AddCommand(elastictranscoder_listJobsByStatusCmd)
}
