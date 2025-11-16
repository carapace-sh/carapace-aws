package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteAccessGrantsLocationCmd = &cobra.Command{
	Use:   "delete-access-grants-location",
	Short: "Deregisters a location from your S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteAccessGrantsLocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deleteAccessGrantsLocationCmd).Standalone()

		s3control_deleteAccessGrantsLocationCmd.Flags().String("access-grants-location-id", "", "The ID of the registered location that you are deregistering from your S3 Access Grants instance.")
		s3control_deleteAccessGrantsLocationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_deleteAccessGrantsLocationCmd.MarkFlagRequired("access-grants-location-id")
		s3control_deleteAccessGrantsLocationCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_deleteAccessGrantsLocationCmd)
}
