package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_createAccessGrantCmd = &cobra.Command{
	Use:   "create-access-grant",
	Short: "Creates an access grant that gives a grantee access to your S3 data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_createAccessGrantCmd).Standalone()

	s3control_createAccessGrantCmd.Flags().String("access-grants-location-configuration", "", "The configuration options of the grant location.")
	s3control_createAccessGrantCmd.Flags().String("access-grants-location-id", "", "The ID of the registered location to which you are granting access.")
	s3control_createAccessGrantCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
	s3control_createAccessGrantCmd.Flags().String("application-arn", "", "The Amazon Resource Name (ARN) of an Amazon Web Services IAM Identity Center application associated with your Identity Center instance.")
	s3control_createAccessGrantCmd.Flags().String("grantee", "", "The user, group, or role to which you are granting access.")
	s3control_createAccessGrantCmd.Flags().String("permission", "", "The type of access that you are granting to your S3 data, which can be set to one of the following values:")
	s3control_createAccessGrantCmd.Flags().String("s3-prefix-type", "", "The type of `S3SubPrefix`.")
	s3control_createAccessGrantCmd.Flags().String("tags", "", "The Amazon Web Services resource tags that you are adding to the access grant.")
	s3control_createAccessGrantCmd.MarkFlagRequired("access-grants-location-id")
	s3control_createAccessGrantCmd.MarkFlagRequired("account-id")
	s3control_createAccessGrantCmd.MarkFlagRequired("grantee")
	s3control_createAccessGrantCmd.MarkFlagRequired("permission")
	s3controlCmd.AddCommand(s3control_createAccessGrantCmd)
}
