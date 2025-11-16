package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeAnomalyDetectorCmd = &cobra.Command{
	Use:   "describe-anomaly-detector",
	Short: "Retrieves detailed information about a specific anomaly detector, including its status and configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeAnomalyDetectorCmd).Standalone()

	amp_describeAnomalyDetectorCmd.Flags().String("anomaly-detector-id", "", "The identifier of the anomaly detector to describe.")
	amp_describeAnomalyDetectorCmd.Flags().String("workspace-id", "", "The identifier of the workspace containing the anomaly detector.")
	amp_describeAnomalyDetectorCmd.MarkFlagRequired("anomaly-detector-id")
	amp_describeAnomalyDetectorCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_describeAnomalyDetectorCmd)
}
