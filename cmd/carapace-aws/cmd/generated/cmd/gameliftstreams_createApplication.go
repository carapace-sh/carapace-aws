package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an application resource in Amazon GameLift Streams, which specifies the application content you want to stream, such as a game build or other software, and configures the settings to run it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_createApplicationCmd).Standalone()

		gameliftstreams_createApplicationCmd.Flags().String("application-log-output-uri", "", "An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs.")
		gameliftstreams_createApplicationCmd.Flags().String("application-log-paths", "", "Locations of log files that your content generates during a stream session.")
		gameliftstreams_createApplicationCmd.Flags().String("application-source-uri", "", "The location of the content that you want to stream.")
		gameliftstreams_createApplicationCmd.Flags().String("client-token", "", "A unique identifier that represents a client request.")
		gameliftstreams_createApplicationCmd.Flags().String("description", "", "A human-readable label for the application.")
		gameliftstreams_createApplicationCmd.Flags().String("executable-path", "", "The relative path and file name of the executable file that Amazon GameLift Streams will stream.")
		gameliftstreams_createApplicationCmd.Flags().String("runtime-environment", "", "Configuration settings that identify the operating system for an application resource.")
		gameliftstreams_createApplicationCmd.Flags().String("tags", "", "A list of labels to assign to the new application resource.")
		gameliftstreams_createApplicationCmd.MarkFlagRequired("application-source-uri")
		gameliftstreams_createApplicationCmd.MarkFlagRequired("description")
		gameliftstreams_createApplicationCmd.MarkFlagRequired("executable-path")
		gameliftstreams_createApplicationCmd.MarkFlagRequired("runtime-environment")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_createApplicationCmd)
}
