package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_copyProductCmd = &cobra.Command{
	Use:   "copy-product",
	Short: "Copies the specified source product to the specified target product or a new product.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_copyProductCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_copyProductCmd).Standalone()

		servicecatalog_copyProductCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_copyProductCmd.Flags().String("copy-options", "", "The copy options.")
		servicecatalog_copyProductCmd.Flags().String("idempotency-token", "", "A unique identifier that you provide to ensure idempotency.")
		servicecatalog_copyProductCmd.Flags().String("source-product-arn", "", "The Amazon Resource Name (ARN) of the source product.")
		servicecatalog_copyProductCmd.Flags().String("source-provisioning-artifact-identifiers", "", "The identifiers of the provisioning artifacts (also known as versions) of the product to copy.")
		servicecatalog_copyProductCmd.Flags().String("target-product-id", "", "The identifier of the target product.")
		servicecatalog_copyProductCmd.Flags().String("target-product-name", "", "A name for the target product.")
		servicecatalog_copyProductCmd.MarkFlagRequired("idempotency-token")
		servicecatalog_copyProductCmd.MarkFlagRequired("source-product-arn")
	})
	servicecatalogCmd.AddCommand(servicecatalog_copyProductCmd)
}
