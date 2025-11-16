package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_listAccessPointsForObjectLambdaCmd = &cobra.Command{
	Use:   "list-access-points-for-object-lambda",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_listAccessPointsForObjectLambdaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_listAccessPointsForObjectLambdaCmd).Standalone()

		s3control_listAccessPointsForObjectLambdaCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified Object Lambda Access Point.")
		s3control_listAccessPointsForObjectLambdaCmd.Flags().String("max-results", "", "The maximum number of access points that you want to include in the list.")
		s3control_listAccessPointsForObjectLambdaCmd.Flags().String("next-token", "", "If the list has more access points than can be returned in one call to this API, this field contains a continuation token that you can provide in subsequent calls to this API to retrieve additional access points.")
		s3control_listAccessPointsForObjectLambdaCmd.MarkFlagRequired("account-id")
	})
	s3controlCmd.AddCommand(s3control_listAccessPointsForObjectLambdaCmd)
}
