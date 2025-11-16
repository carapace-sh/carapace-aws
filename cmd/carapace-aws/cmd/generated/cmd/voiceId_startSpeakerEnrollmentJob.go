package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_startSpeakerEnrollmentJobCmd = &cobra.Command{
	Use:   "start-speaker-enrollment-job",
	Short: "Starts a new batch speaker enrollment job using specified details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_startSpeakerEnrollmentJobCmd).Standalone()

	voiceId_startSpeakerEnrollmentJobCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	voiceId_startSpeakerEnrollmentJobCmd.Flags().String("data-access-role-arn", "", "The IAM role Amazon Resource Name (ARN) that grants Voice ID permissions to access customer's buckets to read the input manifest file and write the job output file.")
	voiceId_startSpeakerEnrollmentJobCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the speaker enrollment job and in which the speakers are enrolled.")
	voiceId_startSpeakerEnrollmentJobCmd.Flags().String("enrollment-config", "", "The enrollment config that contains details such as the action to take when a speaker is already enrolled in Voice ID or when a speaker is identified as a fraudster.")
	voiceId_startSpeakerEnrollmentJobCmd.Flags().String("input-data-config", "", "The input data config containing the S3 location for the input manifest file that contains the list of speaker enrollment requests.")
	voiceId_startSpeakerEnrollmentJobCmd.Flags().String("job-name", "", "A name for your speaker enrollment job.")
	voiceId_startSpeakerEnrollmentJobCmd.Flags().String("output-data-config", "", "The output data config containing the S3 location where Voice ID writes the job output file; you must also include a KMS key ID to encrypt the file.")
	voiceId_startSpeakerEnrollmentJobCmd.MarkFlagRequired("data-access-role-arn")
	voiceId_startSpeakerEnrollmentJobCmd.MarkFlagRequired("domain-id")
	voiceId_startSpeakerEnrollmentJobCmd.MarkFlagRequired("input-data-config")
	voiceId_startSpeakerEnrollmentJobCmd.MarkFlagRequired("output-data-config")
	voiceIdCmd.AddCommand(voiceId_startSpeakerEnrollmentJobCmd)
}
