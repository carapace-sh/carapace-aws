package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_listPipelinesCmd = &cobra.Command{
	Use:   "list-pipelines",
	Short: "The ListPipelines operation gets a list of the pipelines associated with the current AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_listPipelinesCmd).Standalone()

	elastictranscoder_listPipelinesCmd.Flags().String("ascending", "", "To list pipelines in chronological order by the date and time that they were created, enter `true`.")
	elastictranscoder_listPipelinesCmd.Flags().String("page-token", "", "When Elastic Transcoder returns more than one page of results, use `pageToken` in subsequent `GET` requests to get each successive page of results.")
	elastictranscoderCmd.AddCommand(elastictranscoder_listPipelinesCmd)
}
