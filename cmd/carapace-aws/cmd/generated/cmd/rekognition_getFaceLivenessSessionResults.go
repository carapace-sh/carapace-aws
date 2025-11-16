package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_getFaceLivenessSessionResultsCmd = &cobra.Command{
	Use:   "get-face-liveness-session-results",
	Short: "Retrieves the results of a specific Face Liveness session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_getFaceLivenessSessionResultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_getFaceLivenessSessionResultsCmd).Standalone()

		rekognition_getFaceLivenessSessionResultsCmd.Flags().String("session-id", "", "A unique 128-bit UUID.")
		rekognition_getFaceLivenessSessionResultsCmd.MarkFlagRequired("session-id")
	})
	rekognitionCmd.AddCommand(rekognition_getFaceLivenessSessionResultsCmd)
}
