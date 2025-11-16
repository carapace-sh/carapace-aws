package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Permanently deletes an Amazon GameLift Streams application resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_deleteApplicationCmd).Standalone()

	gameliftstreams_deleteApplicationCmd.Flags().String("identifier", "", "An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) or ID that uniquely identifies the application resource.")
	gameliftstreams_deleteApplicationCmd.MarkFlagRequired("identifier")
	gameliftstreamsCmd.AddCommand(gameliftstreams_deleteApplicationCmd)
}
