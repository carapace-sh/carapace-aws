package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Deletes an Amazon EKS cluster control plane.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_deleteClusterCmd).Standalone()

	eks_deleteClusterCmd.Flags().String("name", "", "The name of the cluster to delete.")
	eks_deleteClusterCmd.MarkFlagRequired("name")
	eksCmd.AddCommand(eks_deleteClusterCmd)
}
