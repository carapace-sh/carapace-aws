package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "Lists the Amazon EKS clusters in your Amazon Web Services account in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listClustersCmd).Standalone()

	eks_listClustersCmd.Flags().String("include", "", "Indicates whether external clusters are included in the returned list.")
	eks_listClustersCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
	eks_listClustersCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
	eksCmd.AddCommand(eks_listClustersCmd)
}
