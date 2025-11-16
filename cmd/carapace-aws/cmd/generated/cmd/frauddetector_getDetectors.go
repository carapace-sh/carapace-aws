package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_getDetectorsCmd = &cobra.Command{
	Use:   "get-detectors",
	Short: "Gets all detectors or a single detector if a `detectorId` is specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_getDetectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_getDetectorsCmd).Standalone()

		frauddetector_getDetectorsCmd.Flags().String("detector-id", "", "The detector ID.")
		frauddetector_getDetectorsCmd.Flags().String("max-results", "", "The maximum number of objects to return for the request.")
		frauddetector_getDetectorsCmd.Flags().String("next-token", "", "The next token for the subsequent request.")
	})
	frauddetectorCmd.AddCommand(frauddetector_getDetectorsCmd)
}
