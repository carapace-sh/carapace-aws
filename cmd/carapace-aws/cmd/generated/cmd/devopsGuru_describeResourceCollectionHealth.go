package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeResourceCollectionHealthCmd = &cobra.Command{
	Use:   "describe-resource-collection-health",
	Short: "Returns the number of open proactive insights, open reactive insights, and the Mean Time to Recover (MTTR) for all closed insights in resource collections in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeResourceCollectionHealthCmd).Standalone()

	devopsGuru_describeResourceCollectionHealthCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	devopsGuru_describeResourceCollectionHealthCmd.Flags().String("resource-collection-type", "", "An Amazon Web Services resource collection type.")
	devopsGuru_describeResourceCollectionHealthCmd.MarkFlagRequired("resource-collection-type")
	devopsGuruCmd.AddCommand(devopsGuru_describeResourceCollectionHealthCmd)
}
