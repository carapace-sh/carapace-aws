package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "The *UpdateCluster* API allows you to modify both single-Region and multi-Region cluster configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_updateClusterCmd).Standalone()

	dsql_updateClusterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	dsql_updateClusterCmd.Flags().String("deletion-protection-enabled", "", "Specifies whether to enable deletion protection in your cluster.")
	dsql_updateClusterCmd.Flags().String("identifier", "", "The ID of the cluster you want to update.")
	dsql_updateClusterCmd.Flags().String("kms-encryption-key", "", "The KMS key that encrypts and protects the data on your cluster.")
	dsql_updateClusterCmd.Flags().String("multi-region-properties", "", "The new multi-Region cluster configuration settings to be applied during an update operation.")
	dsql_updateClusterCmd.MarkFlagRequired("identifier")
	dsqlCmd.AddCommand(dsql_updateClusterCmd)
}
