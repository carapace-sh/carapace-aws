package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listTestGridSessionArtifactsCmd = &cobra.Command{
	Use:   "list-test-grid-session-artifacts",
	Short: "Retrieves a list of artifacts created during the session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listTestGridSessionArtifactsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listTestGridSessionArtifactsCmd).Standalone()

		devicefarm_listTestGridSessionArtifactsCmd.Flags().String("max-result", "", "The maximum number of results to be returned by a request.")
		devicefarm_listTestGridSessionArtifactsCmd.Flags().String("next-token", "", "Pagination token.")
		devicefarm_listTestGridSessionArtifactsCmd.Flags().String("session-arn", "", "The ARN of a [TestGridSession]().")
		devicefarm_listTestGridSessionArtifactsCmd.Flags().String("type", "", "Limit results to a specified type of artifact.")
		devicefarm_listTestGridSessionArtifactsCmd.MarkFlagRequired("session-arn")
	})
	devicefarmCmd.AddCommand(devicefarm_listTestGridSessionArtifactsCmd)
}
