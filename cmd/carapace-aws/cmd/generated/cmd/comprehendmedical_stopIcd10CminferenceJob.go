package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_stopIcd10CminferenceJobCmd = &cobra.Command{
	Use:   "stop-icd10-cminference-job",
	Short: "Stops an InferICD10CM inference job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_stopIcd10CminferenceJobCmd).Standalone()

	comprehendmedical_stopIcd10CminferenceJobCmd.Flags().String("job-id", "", "The identifier of the job.")
	comprehendmedical_stopIcd10CminferenceJobCmd.MarkFlagRequired("job-id")
	comprehendmedicalCmd.AddCommand(comprehendmedical_stopIcd10CminferenceJobCmd)
}
