package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_updateStreamGroupCmd = &cobra.Command{
	Use:   "update-stream-group",
	Short: "Updates the configuration settings for an Amazon GameLift Streams stream group resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_updateStreamGroupCmd).Standalone()

	gameliftstreams_updateStreamGroupCmd.Flags().String("default-application-identifier", "", "The unique identifier of the Amazon GameLift Streams application that you want to set as the default application in a stream group.")
	gameliftstreams_updateStreamGroupCmd.Flags().String("description", "", "A descriptive label for the stream group.")
	gameliftstreams_updateStreamGroupCmd.Flags().String("identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the stream group resource.")
	gameliftstreams_updateStreamGroupCmd.Flags().String("location-configurations", "", "A set of one or more locations and the streaming capacity for each location.")
	gameliftstreams_updateStreamGroupCmd.MarkFlagRequired("identifier")
	gameliftstreamsCmd.AddCommand(gameliftstreams_updateStreamGroupCmd)
}
