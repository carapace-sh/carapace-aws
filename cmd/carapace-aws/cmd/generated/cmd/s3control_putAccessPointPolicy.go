package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_putAccessPointPolicyCmd = &cobra.Command{
	Use:   "put-access-point-policy",
	Short: "Associates an access policy with the specified access point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_putAccessPointPolicyCmd).Standalone()

	s3control_putAccessPointPolicyCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for owner of the bucket associated with the specified access point.")
	s3control_putAccessPointPolicyCmd.Flags().String("name", "", "The name of the access point that you want to associate with the specified policy.")
	s3control_putAccessPointPolicyCmd.Flags().String("policy", "", "The policy that you want to apply to the specified access point.")
	s3control_putAccessPointPolicyCmd.MarkFlagRequired("account-id")
	s3control_putAccessPointPolicyCmd.MarkFlagRequired("name")
	s3control_putAccessPointPolicyCmd.MarkFlagRequired("policy")
	s3controlCmd.AddCommand(s3control_putAccessPointPolicyCmd)
}
