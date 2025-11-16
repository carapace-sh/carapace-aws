package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_describeEntitiesDetectionV2JobCmd = &cobra.Command{
	Use:   "describe-entities-detection-v2-job",
	Short: "Gets the properties associated with a medical entities detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_describeEntitiesDetectionV2JobCmd).Standalone()

	comprehendmedical_describeEntitiesDetectionV2JobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend Medical generated for the job.")
	comprehendmedical_describeEntitiesDetectionV2JobCmd.MarkFlagRequired("job-id")
	comprehendmedicalCmd.AddCommand(comprehendmedical_describeEntitiesDetectionV2JobCmd)
}
