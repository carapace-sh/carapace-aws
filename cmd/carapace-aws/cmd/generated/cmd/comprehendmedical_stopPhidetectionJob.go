package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_stopPhidetectionJobCmd = &cobra.Command{
	Use:   "stop-phidetection-job",
	Short: "Stops a protected health information (PHI) detection job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_stopPhidetectionJobCmd).Standalone()

	comprehendmedical_stopPhidetectionJobCmd.Flags().String("job-id", "", "The identifier of the PHI detection job to stop.")
	comprehendmedical_stopPhidetectionJobCmd.MarkFlagRequired("job-id")
	comprehendmedicalCmd.AddCommand(comprehendmedical_stopPhidetectionJobCmd)
}
