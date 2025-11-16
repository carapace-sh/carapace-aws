package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_putClusterPolicyCmd = &cobra.Command{
	Use:   "put-cluster-policy",
	Short: "Attaches a resource-based policy to a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_putClusterPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsql_putClusterPolicyCmd).Standalone()

		dsql_putClusterPolicyCmd.Flags().String("bypass-policy-lockout-safety-check", "", "A flag that allows you to bypass the policy lockout safety check.")
		dsql_putClusterPolicyCmd.Flags().String("client-token", "", "")
		dsql_putClusterPolicyCmd.Flags().String("expected-policy-version", "", "The expected version of the current policy.")
		dsql_putClusterPolicyCmd.Flags().String("identifier", "", "")
		dsql_putClusterPolicyCmd.Flags().String("policy", "", "The resource-based policy document to attach to the cluster.")
		dsql_putClusterPolicyCmd.MarkFlagRequired("identifier")
		dsql_putClusterPolicyCmd.MarkFlagRequired("policy")
	})
	dsqlCmd.AddCommand(dsql_putClusterPolicyCmd)
}
