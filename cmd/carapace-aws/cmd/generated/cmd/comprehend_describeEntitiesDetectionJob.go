package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeEntitiesDetectionJobCmd = &cobra.Command{
	Use:   "describe-entities-detection-job",
	Short: "Gets the properties associated with an entities detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeEntitiesDetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_describeEntitiesDetectionJobCmd).Standalone()

		comprehend_describeEntitiesDetectionJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend generated for the job.")
		comprehend_describeEntitiesDetectionJobCmd.MarkFlagRequired("job-id")
	})
	comprehendCmd.AddCommand(comprehend_describeEntitiesDetectionJobCmd)
}
