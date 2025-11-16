package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_stopEntitiesDetectionJobCmd = &cobra.Command{
	Use:   "stop-entities-detection-job",
	Short: "Stops an entities detection job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_stopEntitiesDetectionJobCmd).Standalone()

	comprehend_stopEntitiesDetectionJobCmd.Flags().String("job-id", "", "The identifier of the entities detection job to stop.")
	comprehend_stopEntitiesDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_stopEntitiesDetectionJobCmd)
}
