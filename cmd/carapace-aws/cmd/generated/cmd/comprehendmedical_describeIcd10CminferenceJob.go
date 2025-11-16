package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_describeIcd10CminferenceJobCmd = &cobra.Command{
	Use:   "describe-icd10-cminference-job",
	Short: "Gets the properties associated with an InferICD10CM job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_describeIcd10CminferenceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_describeIcd10CminferenceJobCmd).Standalone()

		comprehendmedical_describeIcd10CminferenceJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend Medical generated for the job.")
		comprehendmedical_describeIcd10CminferenceJobCmd.MarkFlagRequired("job-id")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_describeIcd10CminferenceJobCmd)
}
