package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listInsightsCmd = &cobra.Command{
	Use:   "list-insights",
	Short: "Returns a list of all insights checked for against the specified cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listInsightsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_listInsightsCmd).Standalone()

		eks_listInsightsCmd.Flags().String("cluster-name", "", "The name of the Amazon EKS cluster associated with the insights.")
		eks_listInsightsCmd.Flags().String("filter", "", "The criteria to filter your list of insights for your cluster.")
		eks_listInsightsCmd.Flags().String("max-results", "", "The maximum number of identity provider configurations returned by `ListInsights` in paginated output.")
		eks_listInsightsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListInsights` request.")
		eks_listInsightsCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_listInsightsCmd)
}
