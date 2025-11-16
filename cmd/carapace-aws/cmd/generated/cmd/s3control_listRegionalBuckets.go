package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listRegionalBucketsCmd = &cobra.Command{
	Use:   "list-regional-buckets",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listRegionalBucketsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_listRegionalBucketsCmd).Standalone()

		s3control_listRegionalBucketsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the Outposts bucket.")
		s3control_listRegionalBucketsCmd.Flags().String("max-results", "", "")
		s3control_listRegionalBucketsCmd.Flags().String("next-token", "", "")
		s3control_listRegionalBucketsCmd.Flags().String("outpost-id", "", "The ID of the Outposts resource.")
		s3control_listRegionalBucketsCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_listRegionalBucketsCmd)
}
