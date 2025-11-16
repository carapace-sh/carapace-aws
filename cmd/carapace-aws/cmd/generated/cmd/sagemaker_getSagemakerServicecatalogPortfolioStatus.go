package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_getSagemakerServicecatalogPortfolioStatusCmd = &cobra.Command{
	Use:   "get-sagemaker-servicecatalog-portfolio-status",
	Short: "Gets the status of Service Catalog in SageMaker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_getSagemakerServicecatalogPortfolioStatusCmd).Standalone()

	sagemakerCmd.AddCommand(sagemaker_getSagemakerServicecatalogPortfolioStatusCmd)
}
