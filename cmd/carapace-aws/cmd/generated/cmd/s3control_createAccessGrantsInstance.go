package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_createAccessGrantsInstanceCmd = &cobra.Command{
	Use:   "create-access-grants-instance",
	Short: "Creates an S3 Access Grants instance, which serves as a logical grouping for access grants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_createAccessGrantsInstanceCmd).Standalone()

	s3control_createAccessGrantsInstanceCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
	s3control_createAccessGrantsInstanceCmd.Flags().String("identity-center-arn", "", "If you would like to associate your S3 Access Grants instance with an Amazon Web Services IAM Identity Center instance, use this field to pass the Amazon Resource Name (ARN) of the Amazon Web Services IAM Identity Center instance that you are associating with your S3 Access Grants instance.")
	s3control_createAccessGrantsInstanceCmd.Flags().String("tags", "", "The Amazon Web Services resource tags that you are adding to the S3 Access Grants instance.")
	s3control_createAccessGrantsInstanceCmd.MarkFlagRequired("account-id")
	s3controlCmd.AddCommand(s3control_createAccessGrantsInstanceCmd)
}
