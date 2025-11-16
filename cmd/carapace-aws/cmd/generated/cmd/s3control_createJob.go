package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "This operation creates an S3 Batch Operations job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_createJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_createJobCmd).Standalone()

		s3control_createJobCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that creates the job.")
		s3control_createJobCmd.Flags().String("client-request-token", "", "An idempotency token to ensure that you don't accidentally submit the same request twice.")
		s3control_createJobCmd.Flags().String("confirmation-required", "", "Indicates whether confirmation is required before Amazon S3 runs the job.")
		s3control_createJobCmd.Flags().String("description", "", "A description for this job.")
		s3control_createJobCmd.Flags().String("manifest", "", "Configuration parameters for the manifest.")
		s3control_createJobCmd.Flags().String("manifest-generator", "", "The attribute container for the ManifestGenerator details.")
		s3control_createJobCmd.Flags().String("operation", "", "The action that you want this job to perform on every object listed in the manifest.")
		s3control_createJobCmd.Flags().String("priority", "", "The numerical priority for this job.")
		s3control_createJobCmd.Flags().String("report", "", "Configuration parameters for the optional job-completion report.")
		s3control_createJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) for the Identity and Access Management (IAM) role that Batch Operations will use to run this job's action on every object in the manifest.")
		s3control_createJobCmd.Flags().String("tags", "", "A set of tags to associate with the S3 Batch Operations job.")
		s3control_createJobCmd.MarkFlagRequired("account-id")
		s3control_createJobCmd.MarkFlagRequired("client-request-token")
		s3control_createJobCmd.MarkFlagRequired("operation")
		s3control_createJobCmd.MarkFlagRequired("priority")
		s3control_createJobCmd.MarkFlagRequired("report")
		s3control_createJobCmd.MarkFlagRequired("role-arn")
	})
	s3controlCmd.AddCommand(s3control_createJobCmd)
}
