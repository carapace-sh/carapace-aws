package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listArtifactsCmd = &cobra.Command{
	Use:   "list-artifacts",
	Short: "Gets information about artifacts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listArtifactsCmd).Standalone()

	devicefarm_listArtifactsCmd.Flags().String("arn", "", "The run, job, suite, or test ARN.")
	devicefarm_listArtifactsCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarm_listArtifactsCmd.Flags().String("type", "", "The artifacts' type.")
	devicefarm_listArtifactsCmd.MarkFlagRequired("arn")
	devicefarm_listArtifactsCmd.MarkFlagRequired("type")
	devicefarmCmd.AddCommand(devicefarm_listArtifactsCmd)
}
