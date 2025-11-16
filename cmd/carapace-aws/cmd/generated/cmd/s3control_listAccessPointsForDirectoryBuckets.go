package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listAccessPointsForDirectoryBucketsCmd = &cobra.Command{
	Use:   "list-access-points-for-directory-buckets",
	Short: "Returns a list of the access points that are owned by the Amazon Web Services account and that are associated with the specified directory bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listAccessPointsForDirectoryBucketsCmd).Standalone()

	s3control_listAccessPointsForDirectoryBucketsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID that owns the access points.")
	s3control_listAccessPointsForDirectoryBucketsCmd.Flags().String("directory-bucket", "", "The name of the directory bucket associated with the access points you want to list.")
	s3control_listAccessPointsForDirectoryBucketsCmd.Flags().String("max-results", "", "The maximum number of access points that you would like returned in the `ListAccessPointsForDirectoryBuckets` response.")
	s3control_listAccessPointsForDirectoryBucketsCmd.Flags().String("next-token", "", "If `NextToken` is returned, there are more access points available than requested in the `maxResults` value.")
	s3control_listAccessPointsForDirectoryBucketsCmd.MarkFlagRequired("account-id")
	s3controlCmd.AddCommand(s3control_listAccessPointsForDirectoryBucketsCmd)
}
