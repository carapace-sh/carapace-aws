package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStackDriftDetectionStatusCmd = &cobra.Command{
	Use:   "describe-stack-drift-detection-status",
	Short: "Returns information about a stack drift detection operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStackDriftDetectionStatusCmd).Standalone()

	cloudformation_describeStackDriftDetectionStatusCmd.Flags().String("stack-drift-detection-id", "", "The ID of the drift detection results of this operation.")
	cloudformation_describeStackDriftDetectionStatusCmd.MarkFlagRequired("stack-drift-detection-id")
	cloudformationCmd.AddCommand(cloudformation_describeStackDriftDetectionStatusCmd)
}
