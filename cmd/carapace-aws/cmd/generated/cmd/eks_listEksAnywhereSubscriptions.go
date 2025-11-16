package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listEksAnywhereSubscriptionsCmd = &cobra.Command{
	Use:   "list-eks-anywhere-subscriptions",
	Short: "Displays the full description of the subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listEksAnywhereSubscriptionsCmd).Standalone()

	eks_listEksAnywhereSubscriptionsCmd.Flags().String("include-status", "", "An array of subscription statuses to filter on.")
	eks_listEksAnywhereSubscriptionsCmd.Flags().String("max-results", "", "The maximum number of cluster results returned by ListEksAnywhereSubscriptions in paginated output.")
	eks_listEksAnywhereSubscriptionsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListEksAnywhereSubscriptions` request where `maxResults` was used and the results exceeded the value of that parameter.")
	eksCmd.AddCommand(eks_listEksAnywhereSubscriptionsCmd)
}
