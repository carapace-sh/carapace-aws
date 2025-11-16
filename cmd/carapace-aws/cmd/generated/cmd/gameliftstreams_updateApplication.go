package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates the mutable configuration settings for a Amazon GameLift Streams application resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_updateApplicationCmd).Standalone()

	gameliftstreams_updateApplicationCmd.Flags().String("application-log-output-uri", "", "An Amazon S3 URI to a bucket where you would like Amazon GameLift Streams to save application logs.")
	gameliftstreams_updateApplicationCmd.Flags().String("application-log-paths", "", "Locations of log files that your content generates during a stream session.")
	gameliftstreams_updateApplicationCmd.Flags().String("description", "", "A human-readable label for the application.")
	gameliftstreams_updateApplicationCmd.Flags().String("identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the application resource.")
	gameliftstreams_updateApplicationCmd.MarkFlagRequired("identifier")
	gameliftstreamsCmd.AddCommand(gameliftstreams_updateApplicationCmd)
}
