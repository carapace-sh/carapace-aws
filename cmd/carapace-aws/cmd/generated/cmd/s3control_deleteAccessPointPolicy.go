package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_deleteAccessPointPolicyCmd = &cobra.Command{
	Use:   "delete-access-point-policy",
	Short: "Deletes the access point policy for the specified access point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_deleteAccessPointPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_deleteAccessPointPolicyCmd).Standalone()

		s3control_deleteAccessPointPolicyCmd.Flags().String("account-id", "", "The account ID for the account that owns the specified access point.")
		s3control_deleteAccessPointPolicyCmd.Flags().String("name", "", "The name of the access point whose policy you want to delete.")
		s3control_deleteAccessPointPolicyCmd.MarkFlagRequired("account-id")
		s3control_deleteAccessPointPolicyCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_deleteAccessPointPolicyCmd)
}
