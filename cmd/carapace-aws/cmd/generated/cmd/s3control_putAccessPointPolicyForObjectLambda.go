package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putAccessPointPolicyForObjectLambdaCmd = &cobra.Command{
	Use:   "put-access-point-policy-for-object-lambda",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putAccessPointPolicyForObjectLambdaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_putAccessPointPolicyForObjectLambdaCmd).Standalone()

		s3control_putAccessPointPolicyForObjectLambdaCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified Object Lambda Access Point.")
		s3control_putAccessPointPolicyForObjectLambdaCmd.Flags().String("name", "", "The name of the Object Lambda Access Point.")
		s3control_putAccessPointPolicyForObjectLambdaCmd.Flags().String("policy", "", "Object Lambda Access Point resource policy document.")
		s3control_putAccessPointPolicyForObjectLambdaCmd.MarkFlagRequired("account-id")
		s3control_putAccessPointPolicyForObjectLambdaCmd.MarkFlagRequired("name")
		s3control_putAccessPointPolicyForObjectLambdaCmd.MarkFlagRequired("policy")
	})
	s3controlCmd.AddCommand(s3control_putAccessPointPolicyForObjectLambdaCmd)
}
