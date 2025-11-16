package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_deregisterClusterCmd = &cobra.Command{
	Use:   "deregister-cluster",
	Short: "Deregisters a connected cluster to remove it from the Amazon EKS control plane.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_deregisterClusterCmd).Standalone()

	eks_deregisterClusterCmd.Flags().String("name", "", "The name of the connected cluster to deregister.")
	eks_deregisterClusterCmd.MarkFlagRequired("name")
	eksCmd.AddCommand(eks_deregisterClusterCmd)
}
