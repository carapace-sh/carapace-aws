package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteS3AccessPolicyCmd = &cobra.Command{
	Use:   "delete-s3-access-policy",
	Short: "Deletes an access policy for the specified store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteS3AccessPolicyCmd).Standalone()

	omics_deleteS3AccessPolicyCmd.Flags().String("s3-access-point-arn", "", "The S3 access point ARN that has the access policy.")
	omics_deleteS3AccessPolicyCmd.MarkFlagRequired("s3-access-point-arn")
	omicsCmd.AddCommand(omics_deleteS3AccessPolicyCmd)
}
