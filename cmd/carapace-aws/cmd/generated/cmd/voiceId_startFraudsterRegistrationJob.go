package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var voiceId_startFraudsterRegistrationJobCmd = &cobra.Command{
	Use:   "start-fraudster-registration-job",
	Short: "Starts a new batch fraudster registration job using provided details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(voiceId_startFraudsterRegistrationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(voiceId_startFraudsterRegistrationJobCmd).Standalone()

		voiceId_startFraudsterRegistrationJobCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		voiceId_startFraudsterRegistrationJobCmd.Flags().String("data-access-role-arn", "", "The IAM role Amazon Resource Name (ARN) that grants Voice ID permissions to access customer's buckets to read the input manifest file and write the Job output file.")
		voiceId_startFraudsterRegistrationJobCmd.Flags().String("domain-id", "", "The identifier of the domain that contains the fraudster registration job and in which the fraudsters are registered.")
		voiceId_startFraudsterRegistrationJobCmd.Flags().String("input-data-config", "", "The input data config containing an S3 URI for the input manifest file that contains the list of fraudster registration requests.")
		voiceId_startFraudsterRegistrationJobCmd.Flags().String("job-name", "", "The name of the new fraudster registration job.")
		voiceId_startFraudsterRegistrationJobCmd.Flags().String("output-data-config", "", "The output data config containing the S3 location where Voice ID writes the job output file; you must also include a KMS key ID to encrypt the file.")
		voiceId_startFraudsterRegistrationJobCmd.Flags().String("registration-config", "", "The registration config containing details such as the action to take when a duplicate fraudster is detected, and the similarity threshold to use for detecting a duplicate fraudster.")
		voiceId_startFraudsterRegistrationJobCmd.MarkFlagRequired("data-access-role-arn")
		voiceId_startFraudsterRegistrationJobCmd.MarkFlagRequired("domain-id")
		voiceId_startFraudsterRegistrationJobCmd.MarkFlagRequired("input-data-config")
		voiceId_startFraudsterRegistrationJobCmd.MarkFlagRequired("output-data-config")
	})
	voiceIdCmd.AddCommand(voiceId_startFraudsterRegistrationJobCmd)
}
