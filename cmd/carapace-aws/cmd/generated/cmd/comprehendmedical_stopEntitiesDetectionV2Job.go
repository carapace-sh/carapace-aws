package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_stopEntitiesDetectionV2JobCmd = &cobra.Command{
	Use:   "stop-entities-detection-v2-job",
	Short: "Stops a medical entities detection job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_stopEntitiesDetectionV2JobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_stopEntitiesDetectionV2JobCmd).Standalone()

		comprehendmedical_stopEntitiesDetectionV2JobCmd.Flags().String("job-id", "", "The identifier of the medical entities job to stop.")
		comprehendmedical_stopEntitiesDetectionV2JobCmd.MarkFlagRequired("job-id")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_stopEntitiesDetectionV2JobCmd)
}
