package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeTopicsDetectionJobCmd = &cobra.Command{
	Use:   "describe-topics-detection-job",
	Short: "Gets the properties associated with a topic detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeTopicsDetectionJobCmd).Standalone()

	comprehend_describeTopicsDetectionJobCmd.Flags().String("job-id", "", "The identifier assigned by the user to the detection job.")
	comprehend_describeTopicsDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_describeTopicsDetectionJobCmd)
}
