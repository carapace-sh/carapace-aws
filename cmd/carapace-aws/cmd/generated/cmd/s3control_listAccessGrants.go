package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listAccessGrantsCmd = &cobra.Command{
	Use:   "list-access-grants",
	Short: "Returns the list of access grants in your S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listAccessGrantsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_listAccessGrantsCmd).Standalone()

		s3control_listAccessGrantsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_listAccessGrantsCmd.Flags().String("application-arn", "", "The Amazon Resource Name (ARN) of an Amazon Web Services IAM Identity Center application associated with your Identity Center instance.")
		s3control_listAccessGrantsCmd.Flags().String("grant-scope", "", "The S3 path of the data to which you are granting access.")
		s3control_listAccessGrantsCmd.Flags().String("grantee-identifier", "", "The unique identifer of the `Grantee`.")
		s3control_listAccessGrantsCmd.Flags().String("grantee-type", "", "The type of the grantee to which access has been granted.")
		s3control_listAccessGrantsCmd.Flags().String("max-results", "", "The maximum number of access grants that you would like returned in the `List Access Grants` response.")
		s3control_listAccessGrantsCmd.Flags().String("next-token", "", "A pagination token to request the next page of results.")
		s3control_listAccessGrantsCmd.Flags().String("permission", "", "The type of permission granted to your S3 data, which can be set to one of the following values:")
		s3control_listAccessGrantsCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_listAccessGrantsCmd)
}
