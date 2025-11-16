package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteAccessGrantCmd = &cobra.Command{
	Use:   "delete-access-grant",
	Short: "Deletes the access grant from the S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteAccessGrantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deleteAccessGrantCmd).Standalone()

		s3control_deleteAccessGrantCmd.Flags().String("access-grant-id", "", "The ID of the access grant.")
		s3control_deleteAccessGrantCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_deleteAccessGrantCmd.MarkFlagRequired("access-grant-id")
		s3control_deleteAccessGrantCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_deleteAccessGrantCmd)
}
