package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listUpdatesCmd = &cobra.Command{
	Use:   "list-updates",
	Short: "Lists the updates associated with an Amazon EKS resource in your Amazon Web Services account, in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listUpdatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_listUpdatesCmd).Standalone()

		eks_listUpdatesCmd.Flags().String("addon-name", "", "The names of the installed add-ons that have available updates.")
		eks_listUpdatesCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
		eks_listUpdatesCmd.Flags().String("name", "", "The name of the Amazon EKS cluster to list updates for.")
		eks_listUpdatesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
		eks_listUpdatesCmd.Flags().String("nodegroup-name", "", "The name of the Amazon EKS managed node group to list updates for.")
		eks_listUpdatesCmd.MarkFlagRequired("name")
	})
	eksCmd.AddCommand(eks_listUpdatesCmd)
}
