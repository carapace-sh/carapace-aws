package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listAccessPointsCmd = &cobra.Command{
	Use:   "list-access-points",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listAccessPointsCmd).Standalone()

	s3control_listAccessPointsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the account that owns the specified access points.")
	s3control_listAccessPointsCmd.Flags().String("bucket", "", "The name of the bucket whose associated access points you want to list.")
	s3control_listAccessPointsCmd.Flags().String("data-source-id", "", "The unique identifier for the data source of the access point.")
	s3control_listAccessPointsCmd.Flags().String("data-source-type", "", "The type of the data source that the access point is attached to.")
	s3control_listAccessPointsCmd.Flags().String("max-results", "", "The maximum number of access points that you want to include in the list.")
	s3control_listAccessPointsCmd.Flags().String("next-token", "", "A continuation token.")
	s3control_listAccessPointsCmd.MarkFlagRequired("account-id")
	s3controlCmd.AddCommand(s3control_listAccessPointsCmd)
}
