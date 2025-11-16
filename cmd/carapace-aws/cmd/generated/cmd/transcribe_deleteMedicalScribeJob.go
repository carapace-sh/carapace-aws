package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_deleteMedicalScribeJobCmd = &cobra.Command{
	Use:   "delete-medical-scribe-job",
	Short: "Deletes a Medical Scribe job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_deleteMedicalScribeJobCmd).Standalone()

	transcribe_deleteMedicalScribeJobCmd.Flags().String("medical-scribe-job-name", "", "The name of the Medical Scribe job you want to delete.")
	transcribe_deleteMedicalScribeJobCmd.MarkFlagRequired("medical-scribe-job-name")
	transcribeCmd.AddCommand(transcribe_deleteMedicalScribeJobCmd)
}
