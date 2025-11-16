package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_startEntitiesDetectionV2JobCmd = &cobra.Command{
	Use:   "start-entities-detection-v2-job",
	Short: "Starts an asynchronous medical entity detection job for a collection of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_startEntitiesDetectionV2JobCmd).Standalone()

	comprehendmedical_startEntitiesDetectionV2JobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
	comprehendmedical_startEntitiesDetectionV2JobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants Amazon Comprehend Medical read access to your input data.")
	comprehendmedical_startEntitiesDetectionV2JobCmd.Flags().String("input-data-config", "", "The input configuration that specifies the format and location of the input data for the job.")
	comprehendmedical_startEntitiesDetectionV2JobCmd.Flags().String("job-name", "", "The identifier of the job.")
	comprehendmedical_startEntitiesDetectionV2JobCmd.Flags().String("kmskey", "", "An AWS Key Management Service key to encrypt your output files.")
	comprehendmedical_startEntitiesDetectionV2JobCmd.Flags().String("language-code", "", "The language of the input documents.")
	comprehendmedical_startEntitiesDetectionV2JobCmd.Flags().String("output-data-config", "", "The output configuration that specifies where to send the output files.")
	comprehendmedical_startEntitiesDetectionV2JobCmd.MarkFlagRequired("data-access-role-arn")
	comprehendmedical_startEntitiesDetectionV2JobCmd.MarkFlagRequired("input-data-config")
	comprehendmedical_startEntitiesDetectionV2JobCmd.MarkFlagRequired("language-code")
	comprehendmedical_startEntitiesDetectionV2JobCmd.MarkFlagRequired("output-data-config")
	comprehendmedicalCmd.AddCommand(comprehendmedical_startEntitiesDetectionV2JobCmd)
}
