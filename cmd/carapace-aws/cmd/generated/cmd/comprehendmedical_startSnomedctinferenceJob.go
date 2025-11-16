package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_startSnomedctinferenceJobCmd = &cobra.Command{
	Use:   "start-snomedctinference-job",
	Short: "Starts an asynchronous job to detect medical concepts and link them to the SNOMED-CT ontology.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_startSnomedctinferenceJobCmd).Standalone()

	comprehendmedical_startSnomedctinferenceJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
	comprehendmedical_startSnomedctinferenceJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants Amazon Comprehend Medical read access to your input data.")
	comprehendmedical_startSnomedctinferenceJobCmd.Flags().String("input-data-config", "", "")
	comprehendmedical_startSnomedctinferenceJobCmd.Flags().String("job-name", "", "The user generated name the asynchronous InferSNOMEDCT job.")
	comprehendmedical_startSnomedctinferenceJobCmd.Flags().String("kmskey", "", "An AWS Key Management Service key used to encrypt your output files.")
	comprehendmedical_startSnomedctinferenceJobCmd.Flags().String("language-code", "", "The language of the input documents.")
	comprehendmedical_startSnomedctinferenceJobCmd.Flags().String("output-data-config", "", "")
	comprehendmedical_startSnomedctinferenceJobCmd.MarkFlagRequired("data-access-role-arn")
	comprehendmedical_startSnomedctinferenceJobCmd.MarkFlagRequired("input-data-config")
	comprehendmedical_startSnomedctinferenceJobCmd.MarkFlagRequired("language-code")
	comprehendmedical_startSnomedctinferenceJobCmd.MarkFlagRequired("output-data-config")
	comprehendmedicalCmd.AddCommand(comprehendmedical_startSnomedctinferenceJobCmd)
}
