package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_listJobsByPipelineCmd = &cobra.Command{
	Use:   "list-jobs-by-pipeline",
	Short: "The ListJobsByPipeline operation gets a list of the jobs currently in a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_listJobsByPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_listJobsByPipelineCmd).Standalone()

		elastictranscoder_listJobsByPipelineCmd.Flags().String("ascending", "", "To list jobs in chronological order by the date and time that they were submitted, enter `true`.")
		elastictranscoder_listJobsByPipelineCmd.Flags().String("page-token", "", "When Elastic Transcoder returns more than one page of results, use `pageToken` in subsequent `GET` requests to get each successive page of results.")
		elastictranscoder_listJobsByPipelineCmd.Flags().String("pipeline-id", "", "The ID of the pipeline for which you want to get job information.")
		elastictranscoder_listJobsByPipelineCmd.MarkFlagRequired("pipeline-id")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_listJobsByPipelineCmd)
}
