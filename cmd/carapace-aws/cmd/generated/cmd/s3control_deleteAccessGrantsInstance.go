package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteAccessGrantsInstanceCmd = &cobra.Command{
	Use:   "delete-access-grants-instance",
	Short: "Deletes your S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteAccessGrantsInstanceCmd).Standalone()

	s3control_deleteAccessGrantsInstanceCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
	s3control_deleteAccessGrantsInstanceCmd.MarkFlagRequired("account-id")
	s3controlCmd.AddCommand(s3control_deleteAccessGrantsInstanceCmd)
}
