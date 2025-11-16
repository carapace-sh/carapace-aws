package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeNodegroupCmd = &cobra.Command{
	Use:   "describe-nodegroup",
	Short: "Describes a managed node group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeNodegroupCmd).Standalone()

	eks_describeNodegroupCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_describeNodegroupCmd.Flags().String("nodegroup-name", "", "The name of the node group to describe.")
	eks_describeNodegroupCmd.MarkFlagRequired("cluster-name")
	eks_describeNodegroupCmd.MarkFlagRequired("nodegroup-name")
	eksCmd.AddCommand(eks_describeNodegroupCmd)
}
