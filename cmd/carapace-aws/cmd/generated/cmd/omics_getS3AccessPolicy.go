package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getS3AccessPolicyCmd = &cobra.Command{
	Use:   "get-s3-access-policy",
	Short: "Retrieves details about an access policy on a given store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getS3AccessPolicyCmd).Standalone()

	omics_getS3AccessPolicyCmd.Flags().String("s3-access-point-arn", "", "The S3 access point ARN that has the access policy.")
	omics_getS3AccessPolicyCmd.MarkFlagRequired("s3-access-point-arn")
	omicsCmd.AddCommand(omics_getS3AccessPolicyCmd)
}
