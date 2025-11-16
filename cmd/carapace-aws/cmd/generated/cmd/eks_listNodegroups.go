package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listNodegroupsCmd = &cobra.Command{
	Use:   "list-nodegroups",
	Short: "Lists the managed node groups associated with the specified cluster in your Amazon Web Services account in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listNodegroupsCmd).Standalone()

	eks_listNodegroupsCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_listNodegroupsCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
	eks_listNodegroupsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
	eks_listNodegroupsCmd.MarkFlagRequired("cluster-name")
	eksCmd.AddCommand(eks_listNodegroupsCmd)
}
