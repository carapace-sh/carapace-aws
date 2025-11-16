package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_deleteAnomalyDetectorCmd = &cobra.Command{
	Use:   "delete-anomaly-detector",
	Short: "Removes an anomaly detector from a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_deleteAnomalyDetectorCmd).Standalone()

	amp_deleteAnomalyDetectorCmd.Flags().String("anomaly-detector-id", "", "The identifier of the anomaly detector to delete.")
	amp_deleteAnomalyDetectorCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	amp_deleteAnomalyDetectorCmd.Flags().String("workspace-id", "", "The identifier of the workspace containing the anomaly detector to delete.")
	amp_deleteAnomalyDetectorCmd.MarkFlagRequired("anomaly-detector-id")
	amp_deleteAnomalyDetectorCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_deleteAnomalyDetectorCmd)
}
