package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_getMedicalScribeJobCmd = &cobra.Command{
	Use:   "get-medical-scribe-job",
	Short: "Provides information about the specified Medical Scribe job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_getMedicalScribeJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_getMedicalScribeJobCmd).Standalone()

		transcribe_getMedicalScribeJobCmd.Flags().String("medical-scribe-job-name", "", "The name of the Medical Scribe job you want information about.")
		transcribe_getMedicalScribeJobCmd.MarkFlagRequired("medical-scribe-job-name")
	})
	transcribeCmd.AddCommand(transcribe_getMedicalScribeJobCmd)
}
