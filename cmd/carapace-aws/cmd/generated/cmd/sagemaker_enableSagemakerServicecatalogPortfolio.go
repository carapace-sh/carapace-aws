package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_enableSagemakerServicecatalogPortfolioCmd = &cobra.Command{
	Use:   "enable-sagemaker-servicecatalog-portfolio",
	Short: "Enables using Service Catalog in SageMaker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_enableSagemakerServicecatalogPortfolioCmd).Standalone()

	sagemakerCmd.AddCommand(sagemaker_enableSagemakerServicecatalogPortfolioCmd)
}
