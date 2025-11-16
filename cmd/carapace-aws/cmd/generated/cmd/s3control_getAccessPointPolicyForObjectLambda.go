package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessPointPolicyForObjectLambdaCmd = &cobra.Command{
	Use:   "get-access-point-policy-for-object-lambda",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessPointPolicyForObjectLambdaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getAccessPointPolicyForObjectLambdaCmd).Standalone()

		s3control_getAccessPointPolicyForObjectLambdaCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified Object Lambda Access Point.")
		s3control_getAccessPointPolicyForObjectLambdaCmd.Flags().String("name", "", "The name of the Object Lambda Access Point.")
		s3control_getAccessPointPolicyForObjectLambdaCmd.MarkFlagRequired("account-id")
		s3control_getAccessPointPolicyForObjectLambdaCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_getAccessPointPolicyForObjectLambdaCmd)
}
