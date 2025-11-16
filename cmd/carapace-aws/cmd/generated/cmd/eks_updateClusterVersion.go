package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_updateClusterVersionCmd = &cobra.Command{
	Use:   "update-cluster-version",
	Short: "Updates an Amazon EKS cluster to the specified Kubernetes version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_updateClusterVersionCmd).Standalone()

	eks_updateClusterVersionCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	eks_updateClusterVersionCmd.Flags().Bool("force", false, "Set this value to `true` to override upgrade-blocking readiness checks when updating a cluster.")
	eks_updateClusterVersionCmd.Flags().String("name", "", "The name of the Amazon EKS cluster to update.")
	eks_updateClusterVersionCmd.Flags().Bool("no-force", false, "Set this value to `true` to override upgrade-blocking readiness checks when updating a cluster.")
	eks_updateClusterVersionCmd.Flags().String("version", "", "The desired Kubernetes version following a successful update.")
	eks_updateClusterVersionCmd.MarkFlagRequired("name")
	eks_updateClusterVersionCmd.Flag("no-force").Hidden = true
	eks_updateClusterVersionCmd.MarkFlagRequired("version")
	eksCmd.AddCommand(eks_updateClusterVersionCmd)
}
