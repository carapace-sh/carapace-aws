package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessGrantsLocationCmd = &cobra.Command{
	Use:   "get-access-grants-location",
	Short: "Retrieves the details of a particular location registered in your S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessGrantsLocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getAccessGrantsLocationCmd).Standalone()

		s3control_getAccessGrantsLocationCmd.Flags().String("access-grants-location-id", "", "The ID of the registered location that you are retrieving.")
		s3control_getAccessGrantsLocationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_getAccessGrantsLocationCmd.MarkFlagRequired("access-grants-location-id")
		s3control_getAccessGrantsLocationCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_getAccessGrantsLocationCmd)
}
