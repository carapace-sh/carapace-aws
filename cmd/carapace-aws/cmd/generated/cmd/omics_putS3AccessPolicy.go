package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_putS3AccessPolicyCmd = &cobra.Command{
	Use:   "put-s3-access-policy",
	Short: "Adds an access policy to the specified store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_putS3AccessPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_putS3AccessPolicyCmd).Standalone()

		omics_putS3AccessPolicyCmd.Flags().String("s3-access-point-arn", "", "The S3 access point ARN where you want to put the access policy.")
		omics_putS3AccessPolicyCmd.Flags().String("s3-access-policy", "", "The resource policy that controls S3 access to the store.")
		omics_putS3AccessPolicyCmd.MarkFlagRequired("s3-access-point-arn")
		omics_putS3AccessPolicyCmd.MarkFlagRequired("s3-access-policy")
	})
	omicsCmd.AddCommand(omics_putS3AccessPolicyCmd)
}
