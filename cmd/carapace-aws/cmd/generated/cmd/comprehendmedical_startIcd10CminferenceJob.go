package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_startIcd10CminferenceJobCmd = &cobra.Command{
	Use:   "start-icd10-cminference-job",
	Short: "Starts an asynchronous job to detect medical conditions and link them to the ICD-10-CM ontology.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_startIcd10CminferenceJobCmd).Standalone()

	comprehendmedical_startIcd10CminferenceJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
	comprehendmedical_startIcd10CminferenceJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants Amazon Comprehend Medical read access to your input data.")
	comprehendmedical_startIcd10CminferenceJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
	comprehendmedical_startIcd10CminferenceJobCmd.Flags().String("job-name", "", "The identifier of the job.")
	comprehendmedical_startIcd10CminferenceJobCmd.Flags().String("kmskey", "", "An AWS Key Management Service key to encrypt your output files.")
	comprehendmedical_startIcd10CminferenceJobCmd.Flags().String("language-code", "", "The language of the input documents.")
	comprehendmedical_startIcd10CminferenceJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
	comprehendmedical_startIcd10CminferenceJobCmd.MarkFlagRequired("data-access-role-arn")
	comprehendmedical_startIcd10CminferenceJobCmd.MarkFlagRequired("input-data-config")
	comprehendmedical_startIcd10CminferenceJobCmd.MarkFlagRequired("language-code")
	comprehendmedical_startIcd10CminferenceJobCmd.MarkFlagRequired("output-data-config")
	comprehendmedicalCmd.AddCommand(comprehendmedical_startIcd10CminferenceJobCmd)
}
