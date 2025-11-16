package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeCopyProductStatusCmd = &cobra.Command{
	Use:   "describe-copy-product-status",
	Short: "Gets the status of the specified copy product operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeCopyProductStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_describeCopyProductStatusCmd).Standalone()

		servicecatalog_describeCopyProductStatusCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_describeCopyProductStatusCmd.Flags().String("copy-product-token", "", "The token for the copy product operation.")
		servicecatalog_describeCopyProductStatusCmd.MarkFlagRequired("copy-product-token")
	})
	servicecatalogCmd.AddCommand(servicecatalog_describeCopyProductStatusCmd)
}
