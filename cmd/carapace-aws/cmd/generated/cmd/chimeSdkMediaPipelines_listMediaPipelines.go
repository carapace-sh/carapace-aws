package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_listMediaPipelinesCmd = &cobra.Command{
	Use:   "list-media-pipelines",
	Short: "Returns a list of media pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_listMediaPipelinesCmd).Standalone()

	chimeSdkMediaPipelines_listMediaPipelinesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chimeSdkMediaPipelines_listMediaPipelinesCmd.Flags().String("next-token", "", "The token used to retrieve the next page of results.")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_listMediaPipelinesCmd)
}
