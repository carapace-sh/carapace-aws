package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeClusterCmd = &cobra.Command{
	Use:   "describe-cluster",
	Short: "Describes an Amazon EKS cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeClusterCmd).Standalone()

	eks_describeClusterCmd.Flags().String("name", "", "The name of your cluster.")
	eks_describeClusterCmd.MarkFlagRequired("name")
	eksCmd.AddCommand(eks_describeClusterCmd)
}
