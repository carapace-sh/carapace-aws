package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteAccessPointPolicyForObjectLambdaCmd = &cobra.Command{
	Use:   "delete-access-point-policy-for-object-lambda",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteAccessPointPolicyForObjectLambdaCmd).Standalone()

	s3control_deleteAccessPointPolicyForObjectLambdaCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified Object Lambda Access Point.")
	s3control_deleteAccessPointPolicyForObjectLambdaCmd.Flags().String("name", "", "The name of the Object Lambda Access Point you want to delete the policy for.")
	s3control_deleteAccessPointPolicyForObjectLambdaCmd.MarkFlagRequired("account-id")
	s3control_deleteAccessPointPolicyForObjectLambdaCmd.MarkFlagRequired("name")
	s3controlCmd.AddCommand(s3control_deleteAccessPointPolicyForObjectLambdaCmd)
}
