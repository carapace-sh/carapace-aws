package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_updateClusterConfigCmd = &cobra.Command{
	Use:   "update-cluster-config",
	Short: "Updates an Amazon EKS cluster configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_updateClusterConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_updateClusterConfigCmd).Standalone()

		eks_updateClusterConfigCmd.Flags().String("access-config", "", "The access configuration for the cluster.")
		eks_updateClusterConfigCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		eks_updateClusterConfigCmd.Flags().String("compute-config", "", "Update the configuration of the compute capability of your EKS Auto Mode cluster.")
		eks_updateClusterConfigCmd.Flags().String("deletion-protection", "", "Specifies whether to enable or disable deletion protection for the cluster.")
		eks_updateClusterConfigCmd.Flags().String("kubernetes-network-config", "", "")
		eks_updateClusterConfigCmd.Flags().String("logging", "", "Enable or disable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs .")
		eks_updateClusterConfigCmd.Flags().String("name", "", "The name of the Amazon EKS cluster to update.")
		eks_updateClusterConfigCmd.Flags().String("remote-network-config", "", "")
		eks_updateClusterConfigCmd.Flags().String("resources-vpc-config", "", "")
		eks_updateClusterConfigCmd.Flags().String("storage-config", "", "Update the configuration of the block storage capability of your EKS Auto Mode cluster.")
		eks_updateClusterConfigCmd.Flags().String("upgrade-policy", "", "You can enable or disable extended support for clusters currently on standard support.")
		eks_updateClusterConfigCmd.Flags().String("zonal-shift-config", "", "Enable or disable ARC zonal shift for the cluster.")
		eks_updateClusterConfigCmd.MarkFlagRequired("name")
	})
	eksCmd.AddCommand(eks_updateClusterConfigCmd)
}
