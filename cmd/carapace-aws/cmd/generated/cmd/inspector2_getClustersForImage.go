package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getClustersForImageCmd = &cobra.Command{
	Use:   "get-clusters-for-image",
	Short: "Returns a list of clusters and metadata associated with an image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getClustersForImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_getClustersForImageCmd).Standalone()

		inspector2_getClustersForImageCmd.Flags().String("filter", "", "The resource Id for the Amazon ECR image.")
		inspector2_getClustersForImageCmd.Flags().String("max-results", "", "The maximum number of results to be returned in a single page of results.")
		inspector2_getClustersForImageCmd.Flags().String("next-token", "", "The pagination token from a previous request used to retrieve the next page of results.")
		inspector2_getClustersForImageCmd.MarkFlagRequired("filter")
	})
	inspector2Cmd.AddCommand(inspector2_getClustersForImageCmd)
}
