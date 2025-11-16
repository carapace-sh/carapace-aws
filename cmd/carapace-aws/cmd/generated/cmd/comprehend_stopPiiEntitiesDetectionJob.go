package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_stopPiiEntitiesDetectionJobCmd = &cobra.Command{
	Use:   "stop-pii-entities-detection-job",
	Short: "Stops a PII entities detection job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_stopPiiEntitiesDetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_stopPiiEntitiesDetectionJobCmd).Standalone()

		comprehend_stopPiiEntitiesDetectionJobCmd.Flags().String("job-id", "", "The identifier of the PII entities detection job to stop.")
		comprehend_stopPiiEntitiesDetectionJobCmd.MarkFlagRequired("job-id")
	})
	comprehendCmd.AddCommand(comprehend_stopPiiEntitiesDetectionJobCmd)
}
