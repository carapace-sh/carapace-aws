package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listAddonsCmd = &cobra.Command{
	Use:   "list-addons",
	Short: "Lists the installed add-ons.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listAddonsCmd).Standalone()

	eks_listAddonsCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_listAddonsCmd.Flags().String("max-results", "", "The maximum number of results, returned in paginated output.")
	eks_listAddonsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated request, where `maxResults` was used and the results exceeded the value of that parameter.")
	eks_listAddonsCmd.MarkFlagRequired("cluster-name")
	eksCmd.AddCommand(eks_listAddonsCmd)
}
