package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_disableSagemakerServicecatalogPortfolioCmd = &cobra.Command{
	Use:   "disable-sagemaker-servicecatalog-portfolio",
	Short: "Disables using Service Catalog in SageMaker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_disableSagemakerServicecatalogPortfolioCmd).Standalone()

	sagemakerCmd.AddCommand(sagemaker_disableSagemakerServicecatalogPortfolioCmd)
}
