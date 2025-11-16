package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_stopSnomedctinferenceJobCmd = &cobra.Command{
	Use:   "stop-snomedctinference-job",
	Short: "Stops an InferSNOMEDCT inference job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_stopSnomedctinferenceJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_stopSnomedctinferenceJobCmd).Standalone()

		comprehendmedical_stopSnomedctinferenceJobCmd.Flags().String("job-id", "", "The job id of the asynchronous InferSNOMEDCT job to be stopped.")
		comprehendmedical_stopSnomedctinferenceJobCmd.MarkFlagRequired("job-id")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_stopSnomedctinferenceJobCmd)
}
