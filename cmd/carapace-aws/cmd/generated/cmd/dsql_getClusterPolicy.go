package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_getClusterPolicyCmd = &cobra.Command{
	Use:   "get-cluster-policy",
	Short: "Retrieves the resource-based policy document attached to a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_getClusterPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsql_getClusterPolicyCmd).Standalone()

		dsql_getClusterPolicyCmd.Flags().String("identifier", "", "The ID of the cluster to retrieve the policy from.")
		dsql_getClusterPolicyCmd.MarkFlagRequired("identifier")
	})
	dsqlCmd.AddCommand(dsql_getClusterPolicyCmd)
}
