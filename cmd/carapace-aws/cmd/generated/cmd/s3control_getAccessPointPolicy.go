package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessPointPolicyCmd = &cobra.Command{
	Use:   "get-access-point-policy",
	Short: "Returns the access point policy associated with the specified access point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessPointPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getAccessPointPolicyCmd).Standalone()

		s3control_getAccessPointPolicyCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified access point.")
		s3control_getAccessPointPolicyCmd.Flags().String("name", "", "The name of the access point whose policy you want to retrieve.")
		s3control_getAccessPointPolicyCmd.MarkFlagRequired("account-id")
		s3control_getAccessPointPolicyCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_getAccessPointPolicyCmd)
}
