package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_listMediaCapturePipelinesCmd = &cobra.Command{
	Use:   "list-media-capture-pipelines",
	Short: "Returns a list of media pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_listMediaCapturePipelinesCmd).Standalone()

	chimeSdkMediaPipelines_listMediaCapturePipelinesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chimeSdkMediaPipelines_listMediaCapturePipelinesCmd.Flags().String("next-token", "", "The token used to retrieve the next page of results.")
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_listMediaCapturePipelinesCmd)
}
