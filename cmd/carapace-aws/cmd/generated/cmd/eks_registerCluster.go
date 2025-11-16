package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_registerClusterCmd = &cobra.Command{
	Use:   "register-cluster",
	Short: "Connects a Kubernetes cluster to the Amazon EKS control plane.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_registerClusterCmd).Standalone()

	eks_registerClusterCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	eks_registerClusterCmd.Flags().String("connector-config", "", "The configuration settings required to connect the Kubernetes cluster to the Amazon EKS control plane.")
	eks_registerClusterCmd.Flags().String("name", "", "A unique name for this cluster in your Amazon Web Services Region.")
	eks_registerClusterCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
	eks_registerClusterCmd.MarkFlagRequired("connector-config")
	eks_registerClusterCmd.MarkFlagRequired("name")
	eksCmd.AddCommand(eks_registerClusterCmd)
}
