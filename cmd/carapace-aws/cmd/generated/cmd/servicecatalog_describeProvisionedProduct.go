package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeProvisionedProductCmd = &cobra.Command{
	Use:   "describe-provisioned-product",
	Short: "Gets information about the specified provisioned product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeProvisionedProductCmd).Standalone()

	servicecatalog_describeProvisionedProductCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_describeProvisionedProductCmd.Flags().String("id", "", "The provisioned product identifier.")
	servicecatalog_describeProvisionedProductCmd.Flags().String("name", "", "The name of the provisioned product.")
	servicecatalogCmd.AddCommand(servicecatalog_describeProvisionedProductCmd)
}
