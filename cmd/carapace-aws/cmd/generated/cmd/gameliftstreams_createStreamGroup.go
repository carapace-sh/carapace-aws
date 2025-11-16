package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_createStreamGroupCmd = &cobra.Command{
	Use:   "create-stream-group",
	Short: "Stream groups manage how Amazon GameLift Streams allocates resources and handles concurrent streams, allowing you to effectively manage capacity and costs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_createStreamGroupCmd).Standalone()

	gameliftstreams_createStreamGroupCmd.Flags().String("client-token", "", "A unique identifier that represents a client request.")
	gameliftstreams_createStreamGroupCmd.Flags().String("default-application-identifier", "", "The unique identifier of the Amazon GameLift Streams application that you want to set as the default application in a stream group.")
	gameliftstreams_createStreamGroupCmd.Flags().String("description", "", "A descriptive label for the stream group.")
	gameliftstreams_createStreamGroupCmd.Flags().String("location-configurations", "", "A set of one or more locations and the streaming capacity for each location.")
	gameliftstreams_createStreamGroupCmd.Flags().String("stream-class", "", "The target stream quality for sessions that are hosted in this stream group.")
	gameliftstreams_createStreamGroupCmd.Flags().String("tags", "", "A list of labels to assign to the new stream group resource.")
	gameliftstreams_createStreamGroupCmd.MarkFlagRequired("description")
	gameliftstreams_createStreamGroupCmd.MarkFlagRequired("stream-class")
	gameliftstreamsCmd.AddCommand(gameliftstreams_createStreamGroupCmd)
}
