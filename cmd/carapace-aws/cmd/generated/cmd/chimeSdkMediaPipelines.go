package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelinesCmd = &cobra.Command{
	Use:   "chime-sdk-media-pipelines",
	Short: "The Amazon Chime SDK media pipeline APIs in this section allow software developers to create Amazon Chime SDK media pipelines that capture, concatenate, or stream your Amazon Chime SDK meetings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelinesCmd).Standalone()

	rootCmd.AddCommand(chimeSdkMediaPipelinesCmd)
}
