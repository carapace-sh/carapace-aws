package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_getStreamGroupCmd = &cobra.Command{
	Use:   "get-stream-group",
	Short: "Retrieves properties for a Amazon GameLift Streams stream group resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_getStreamGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_getStreamGroupCmd).Standalone()

		gameliftstreams_getStreamGroupCmd.Flags().String("identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream group resource.")
		gameliftstreams_getStreamGroupCmd.MarkFlagRequired("identifier")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_getStreamGroupCmd)
}
