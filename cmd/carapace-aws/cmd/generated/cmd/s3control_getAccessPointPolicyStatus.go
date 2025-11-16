package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessPointPolicyStatusCmd = &cobra.Command{
	Use:   "get-access-point-policy-status",
	Short: "This operation is not supported by directory buckets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessPointPolicyStatusCmd).Standalone()

	s3control_getAccessPointPolicyStatusCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified access point.")
	s3control_getAccessPointPolicyStatusCmd.Flags().String("name", "", "The name of the access point whose policy status you want to retrieve.")
	s3control_getAccessPointPolicyStatusCmd.MarkFlagRequired("account-id")
	s3control_getAccessPointPolicyStatusCmd.MarkFlagRequired("name")
	s3controlCmd.AddCommand(s3control_getAccessPointPolicyStatusCmd)
}
