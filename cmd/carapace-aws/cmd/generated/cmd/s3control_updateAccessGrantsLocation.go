package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_updateAccessGrantsLocationCmd = &cobra.Command{
	Use:   "update-access-grants-location",
	Short: "Updates the IAM role of a registered location in your S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_updateAccessGrantsLocationCmd).Standalone()

	s3control_updateAccessGrantsLocationCmd.Flags().String("access-grants-location-id", "", "The ID of the registered location that you are updating.")
	s3control_updateAccessGrantsLocationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
	s3control_updateAccessGrantsLocationCmd.Flags().String("iamrole-arn", "", "The Amazon Resource Name (ARN) of the IAM role for the registered location.")
	s3control_updateAccessGrantsLocationCmd.MarkFlagRequired("access-grants-location-id")
	s3control_updateAccessGrantsLocationCmd.MarkFlagRequired("account-id")
	s3control_updateAccessGrantsLocationCmd.MarkFlagRequired("iamrole-arn")
	s3controlCmd.AddCommand(s3control_updateAccessGrantsLocationCmd)
}
