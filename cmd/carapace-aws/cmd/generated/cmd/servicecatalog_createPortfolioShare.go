package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_createPortfolioShareCmd = &cobra.Command{
	Use:   "create-portfolio-share",
	Short: "Shares the specified portfolio with the specified account or organization node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_createPortfolioShareCmd).Standalone()

	servicecatalog_createPortfolioShareCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_createPortfolioShareCmd.Flags().String("account-id", "", "The Amazon Web Services account ID.")
	servicecatalog_createPortfolioShareCmd.Flags().Bool("no-share-principals", false, "This parameter is only supported for portfolios with an **OrganizationalNode** Type of `ORGANIZATION` or `ORGANIZATIONAL_UNIT`.")
	servicecatalog_createPortfolioShareCmd.Flags().Bool("no-share-tag-options", false, "Enables or disables `TagOptions` sharing when creating the portfolio share.")
	servicecatalog_createPortfolioShareCmd.Flags().String("organization-node", "", "The organization node to whom you are going to share.")
	servicecatalog_createPortfolioShareCmd.Flags().String("portfolio-id", "", "The portfolio identifier.")
	servicecatalog_createPortfolioShareCmd.Flags().Bool("share-principals", false, "This parameter is only supported for portfolios with an **OrganizationalNode** Type of `ORGANIZATION` or `ORGANIZATIONAL_UNIT`.")
	servicecatalog_createPortfolioShareCmd.Flags().Bool("share-tag-options", false, "Enables or disables `TagOptions` sharing when creating the portfolio share.")
	servicecatalog_createPortfolioShareCmd.Flag("no-share-principals").Hidden = true
	servicecatalog_createPortfolioShareCmd.Flag("no-share-tag-options").Hidden = true
	servicecatalog_createPortfolioShareCmd.MarkFlagRequired("portfolio-id")
	servicecatalogCmd.AddCommand(servicecatalog_createPortfolioShareCmd)
}
