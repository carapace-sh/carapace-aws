package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_createAccessGrantsLocationCmd = &cobra.Command{
	Use:   "create-access-grants-location",
	Short: "The S3 data location that you would like to register in your S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_createAccessGrantsLocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_createAccessGrantsLocationCmd).Standalone()

		s3control_createAccessGrantsLocationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_createAccessGrantsLocationCmd.Flags().String("iamrole-arn", "", "The Amazon Resource Name (ARN) of the IAM role for the registered location.")
		s3control_createAccessGrantsLocationCmd.Flags().String("location-scope", "", "The S3 path to the location that you are registering.")
		s3control_createAccessGrantsLocationCmd.Flags().String("tags", "", "The Amazon Web Services resource tags that you are adding to the S3 Access Grants location.")
		s3control_createAccessGrantsLocationCmd.MarkFlagRequired("account-id")
		s3control_createAccessGrantsLocationCmd.MarkFlagRequired("iamrole-arn")
		s3control_createAccessGrantsLocationCmd.MarkFlagRequired("location-scope")
	})
	s3controlCmd.AddCommand(s3control_createAccessGrantsLocationCmd)
}
