package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_describeSnomedctinferenceJobCmd = &cobra.Command{
	Use:   "describe-snomedctinference-job",
	Short: "Gets the properties associated with an InferSNOMEDCT job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_describeSnomedctinferenceJobCmd).Standalone()

	comprehendmedical_describeSnomedctinferenceJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend Medical generated for the job.")
	comprehendmedical_describeSnomedctinferenceJobCmd.MarkFlagRequired("job-id")
	comprehendmedicalCmd.AddCommand(comprehendmedical_describeSnomedctinferenceJobCmd)
}
