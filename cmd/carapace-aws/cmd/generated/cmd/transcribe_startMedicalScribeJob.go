package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_startMedicalScribeJobCmd = &cobra.Command{
	Use:   "start-medical-scribe-job",
	Short: "Transcribes patient-clinician conversations and generates clinical notes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_startMedicalScribeJobCmd).Standalone()

	transcribe_startMedicalScribeJobCmd.Flags().String("channel-definitions", "", "Makes it possible to specify which speaker is on which channel.")
	transcribe_startMedicalScribeJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files, write to the output bucket, and use your KMS key if supplied.")
	transcribe_startMedicalScribeJobCmd.Flags().String("kmsencryption-context", "", "A map of plain text, non-secret key:value pairs, known as encryption context pairs, that provide an added layer of security for your data.")
	transcribe_startMedicalScribeJobCmd.Flags().String("media", "", "")
	transcribe_startMedicalScribeJobCmd.Flags().String("medical-scribe-context", "", "The `MedicalScribeContext` object that contains contextual information which is used during clinical note generation to add relevant context to the note.")
	transcribe_startMedicalScribeJobCmd.Flags().String("medical-scribe-job-name", "", "A unique name, chosen by you, for your Medical Scribe job.")
	transcribe_startMedicalScribeJobCmd.Flags().String("output-bucket-name", "", "The name of the Amazon S3 bucket where you want your Medical Scribe output stored.")
	transcribe_startMedicalScribeJobCmd.Flags().String("output-encryption-kmskey-id", "", "The Amazon Resource Name (ARN) of a KMS key that you want to use to encrypt your Medical Scribe output.")
	transcribe_startMedicalScribeJobCmd.Flags().String("settings", "", "Makes it possible to control how your Medical Scribe job is processed using a `MedicalScribeSettings` object.")
	transcribe_startMedicalScribeJobCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to the Medical Scribe job.")
	transcribe_startMedicalScribeJobCmd.MarkFlagRequired("data-access-role-arn")
	transcribe_startMedicalScribeJobCmd.MarkFlagRequired("media")
	transcribe_startMedicalScribeJobCmd.MarkFlagRequired("medical-scribe-job-name")
	transcribe_startMedicalScribeJobCmd.MarkFlagRequired("output-bucket-name")
	transcribe_startMedicalScribeJobCmd.MarkFlagRequired("settings")
	transcribeCmd.AddCommand(transcribe_startMedicalScribeJobCmd)
}
