package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_startRxNormInferenceJobCmd = &cobra.Command{
	Use:   "start-rx-norm-inference-job",
	Short: "Starts an asynchronous job to detect medication entities and link them to the RxNorm ontology.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_startRxNormInferenceJobCmd).Standalone()

	comprehendmedical_startRxNormInferenceJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
	comprehendmedical_startRxNormInferenceJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants Amazon Comprehend Medical read access to your input data.")
	comprehendmedical_startRxNormInferenceJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
	comprehendmedical_startRxNormInferenceJobCmd.Flags().String("job-name", "", "The identifier of the job.")
	comprehendmedical_startRxNormInferenceJobCmd.Flags().String("kmskey", "", "An AWS Key Management Service key to encrypt your output files.")
	comprehendmedical_startRxNormInferenceJobCmd.Flags().String("language-code", "", "The language of the input documents.")
	comprehendmedical_startRxNormInferenceJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
	comprehendmedical_startRxNormInferenceJobCmd.MarkFlagRequired("data-access-role-arn")
	comprehendmedical_startRxNormInferenceJobCmd.MarkFlagRequired("input-data-config")
	comprehendmedical_startRxNormInferenceJobCmd.MarkFlagRequired("language-code")
	comprehendmedical_startRxNormInferenceJobCmd.MarkFlagRequired("output-data-config")
	comprehendmedicalCmd.AddCommand(comprehendmedical_startRxNormInferenceJobCmd)
}
