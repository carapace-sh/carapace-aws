package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_deleteClusterPolicyCmd = &cobra.Command{
	Use:   "delete-cluster-policy",
	Short: "Deletes the resource-based policy attached to a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_deleteClusterPolicyCmd).Standalone()

	dsql_deleteClusterPolicyCmd.Flags().String("client-token", "", "")
	dsql_deleteClusterPolicyCmd.Flags().String("expected-policy-version", "", "The expected version of the policy to delete.")
	dsql_deleteClusterPolicyCmd.Flags().String("identifier", "", "")
	dsql_deleteClusterPolicyCmd.MarkFlagRequired("identifier")
	dsqlCmd.AddCommand(dsql_deleteClusterPolicyCmd)
}
