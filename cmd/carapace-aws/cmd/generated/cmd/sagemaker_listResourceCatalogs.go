package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listResourceCatalogsCmd = &cobra.Command{
	Use:   "list-resource-catalogs",
	Short: "Lists Amazon SageMaker Catalogs based on given filters and orders.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listResourceCatalogsCmd).Standalone()

	sagemaker_listResourceCatalogsCmd.Flags().String("creation-time-after", "", "Use this parameter to search for `ResourceCatalog`s created after a specific date and time.")
	sagemaker_listResourceCatalogsCmd.Flags().String("creation-time-before", "", "Use this parameter to search for `ResourceCatalog`s created before a specific date and time.")
	sagemaker_listResourceCatalogsCmd.Flags().String("max-results", "", "The maximum number of results returned by `ListResourceCatalogs`.")
	sagemaker_listResourceCatalogsCmd.Flags().String("name-contains", "", "A string that partially matches one or more `ResourceCatalog`s names.")
	sagemaker_listResourceCatalogsCmd.Flags().String("next-token", "", "A token to resume pagination of `ListResourceCatalogs` results.")
	sagemaker_listResourceCatalogsCmd.Flags().String("sort-by", "", "The value on which the resource catalog list is sorted.")
	sagemaker_listResourceCatalogsCmd.Flags().String("sort-order", "", "The order in which the resource catalogs are listed.")
	sagemakerCmd.AddCommand(sagemaker_listResourceCatalogsCmd)
}
