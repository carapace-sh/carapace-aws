package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_deleteStreamGroupCmd = &cobra.Command{
	Use:   "delete-stream-group",
	Short: "Permanently deletes all compute resources and information related to a stream group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_deleteStreamGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_deleteStreamGroupCmd).Standalone()

		gameliftstreams_deleteStreamGroupCmd.Flags().String("identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream group resource.")
		gameliftstreams_deleteStreamGroupCmd.MarkFlagRequired("identifier")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_deleteStreamGroupCmd)
}
