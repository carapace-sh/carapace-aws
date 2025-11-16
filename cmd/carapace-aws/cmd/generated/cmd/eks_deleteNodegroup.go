package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_deleteNodegroupCmd = &cobra.Command{
	Use:   "delete-nodegroup",
	Short: "Deletes a managed node group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_deleteNodegroupCmd).Standalone()

	eks_deleteNodegroupCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_deleteNodegroupCmd.Flags().String("nodegroup-name", "", "The name of the node group to delete.")
	eks_deleteNodegroupCmd.MarkFlagRequired("cluster-name")
	eks_deleteNodegroupCmd.MarkFlagRequired("nodegroup-name")
	eksCmd.AddCommand(eks_deleteNodegroupCmd)
}
