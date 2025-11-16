package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessGrantCmd = &cobra.Command{
	Use:   "get-access-grant",
	Short: "Get the details of an access grant from your S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessGrantCmd).Standalone()

	s3control_getAccessGrantCmd.Flags().String("access-grant-id", "", "The ID of the access grant.")
	s3control_getAccessGrantCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
	s3control_getAccessGrantCmd.MarkFlagRequired("access-grant-id")
	s3control_getAccessGrantCmd.MarkFlagRequired("account-id")
	s3controlCmd.AddCommand(s3control_getAccessGrantCmd)
}
