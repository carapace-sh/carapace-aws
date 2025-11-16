package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Retrieves properties for an Amazon GameLift Streams application resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_getApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_getApplicationCmd).Standalone()

		gameliftstreams_getApplicationCmd.Flags().String("identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the application resource.")
		gameliftstreams_getApplicationCmd.MarkFlagRequired("identifier")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_getApplicationCmd)
}
