package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "The CreateCluster API allows you to create both single-Region clusters and multi-Region clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_createClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsql_createClusterCmd).Standalone()

		dsql_createClusterCmd.Flags().String("bypass-policy-lockout-safety-check", "", "An optional field that controls whether to bypass the lockout prevention check.")
		dsql_createClusterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		dsql_createClusterCmd.Flags().String("deletion-protection-enabled", "", "If enabled, you can't delete your cluster.")
		dsql_createClusterCmd.Flags().String("kms-encryption-key", "", "The KMS key that encrypts and protects the data on your cluster.")
		dsql_createClusterCmd.Flags().String("multi-region-properties", "", "The configuration settings when creating a multi-Region cluster, including the witness region and linked cluster properties.")
		dsql_createClusterCmd.Flags().String("policy", "", "An optional resource-based policy document in JSON format that defines access permissions for the cluster.")
		dsql_createClusterCmd.Flags().String("tags", "", "A map of key and value pairs to use to tag your cluster.")
	})
	dsqlCmd.AddCommand(dsql_createClusterCmd)
}
