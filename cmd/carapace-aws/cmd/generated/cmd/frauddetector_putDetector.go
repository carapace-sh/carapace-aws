package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_putDetectorCmd = &cobra.Command{
	Use:   "put-detector",
	Short: "Creates or updates a detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_putDetectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_putDetectorCmd).Standalone()

		frauddetector_putDetectorCmd.Flags().String("description", "", "The description of the detector.")
		frauddetector_putDetectorCmd.Flags().String("detector-id", "", "The detector ID.")
		frauddetector_putDetectorCmd.Flags().String("event-type-name", "", "The name of the event type.")
		frauddetector_putDetectorCmd.Flags().String("tags", "", "A collection of key and value pairs.")
		frauddetector_putDetectorCmd.MarkFlagRequired("detector-id")
		frauddetector_putDetectorCmd.MarkFlagRequired("event-type-name")
	})
	frauddetectorCmd.AddCommand(frauddetector_putDetectorCmd)
}
