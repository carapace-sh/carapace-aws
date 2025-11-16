package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeClusterVersionsCmd = &cobra.Command{
	Use:   "describe-cluster-versions",
	Short: "Lists available Kubernetes versions for Amazon EKS clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeClusterVersionsCmd).Standalone()

	eks_describeClusterVersionsCmd.Flags().String("cluster-type", "", "The type of cluster to filter versions by.")
	eks_describeClusterVersionsCmd.Flags().String("cluster-versions", "", "List of specific cluster versions to describe.")
	eks_describeClusterVersionsCmd.Flags().String("default-only", "", "Filter to show only default versions.")
	eks_describeClusterVersionsCmd.Flags().String("include-all", "", "Include all available versions in the response.")
	eks_describeClusterVersionsCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	eks_describeClusterVersionsCmd.Flags().String("next-token", "", "Pagination token for the next set of results.")
	eks_describeClusterVersionsCmd.Flags().String("status", "", "This field is deprecated.")
	eks_describeClusterVersionsCmd.Flags().String("version-status", "", "Filter versions by their current status.")
	eksCmd.AddCommand(eks_describeClusterVersionsCmd)
}
