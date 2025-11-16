package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describePiiEntitiesDetectionJobCmd = &cobra.Command{
	Use:   "describe-pii-entities-detection-job",
	Short: "Gets the properties associated with a PII entities detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describePiiEntitiesDetectionJobCmd).Standalone()

	comprehend_describePiiEntitiesDetectionJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend generated for the job.")
	comprehend_describePiiEntitiesDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_describePiiEntitiesDetectionJobCmd)
}
