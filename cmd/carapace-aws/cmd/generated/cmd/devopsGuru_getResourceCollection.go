package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_getResourceCollectionCmd = &cobra.Command{
	Use:   "get-resource-collection",
	Short: "Returns lists Amazon Web Services resources that are of the specified resource collection type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_getResourceCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_getResourceCollectionCmd).Standalone()

		devopsGuru_getResourceCollectionCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
		devopsGuru_getResourceCollectionCmd.Flags().String("resource-collection-type", "", "The type of Amazon Web Services resource collections to return.")
		devopsGuru_getResourceCollectionCmd.MarkFlagRequired("resource-collection-type")
	})
	devopsGuruCmd.AddCommand(devopsGuru_getResourceCollectionCmd)
}
