package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_associateAccessGrantsIdentityCenterCmd = &cobra.Command{
	Use:   "associate-access-grants-identity-center",
	Short: "Associate your S3 Access Grants instance with an Amazon Web Services IAM Identity Center instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_associateAccessGrantsIdentityCenterCmd).Standalone()

	s3control_associateAccessGrantsIdentityCenterCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
	s3control_associateAccessGrantsIdentityCenterCmd.Flags().String("identity-center-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services IAM Identity Center instance that you are associating with your S3 Access Grants instance.")
	s3control_associateAccessGrantsIdentityCenterCmd.MarkFlagRequired("account-id")
	s3control_associateAccessGrantsIdentityCenterCmd.MarkFlagRequired("identity-center-arn")
	s3controlCmd.AddCommand(s3control_associateAccessGrantsIdentityCenterCmd)
}
