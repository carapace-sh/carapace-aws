package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listAccessGrantsLocationsCmd = &cobra.Command{
	Use:   "list-access-grants-locations",
	Short: "Returns a list of the locations registered in your S3 Access Grants instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listAccessGrantsLocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_listAccessGrantsLocationsCmd).Standalone()

		s3control_listAccessGrantsLocationsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the S3 Access Grants instance.")
		s3control_listAccessGrantsLocationsCmd.Flags().String("location-scope", "", "The S3 path to the location that you are registering.")
		s3control_listAccessGrantsLocationsCmd.Flags().String("max-results", "", "The maximum number of access grants that you would like returned in the `List Access Grants` response.")
		s3control_listAccessGrantsLocationsCmd.Flags().String("next-token", "", "A pagination token to request the next page of results.")
		s3control_listAccessGrantsLocationsCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_listAccessGrantsLocationsCmd)
}
