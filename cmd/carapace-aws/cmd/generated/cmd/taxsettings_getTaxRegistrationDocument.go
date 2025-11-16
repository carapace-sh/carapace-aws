package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var taxsettings_getTaxRegistrationDocumentCmd = &cobra.Command{
	Use:   "get-tax-registration-document",
	Short: "Downloads your tax documents to the Amazon S3 bucket that you specify in your request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taxsettings_getTaxRegistrationDocumentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(taxsettings_getTaxRegistrationDocumentCmd).Standalone()

		taxsettings_getTaxRegistrationDocumentCmd.Flags().String("destination-s3-location", "", "The Amazon S3 bucket that you specify to download your tax documents to.")
		taxsettings_getTaxRegistrationDocumentCmd.Flags().String("tax-document-metadata", "", "The metadata for your tax document.")
		taxsettings_getTaxRegistrationDocumentCmd.MarkFlagRequired("tax-document-metadata")
	})
	taxsettingsCmd.AddCommand(taxsettings_getTaxRegistrationDocumentCmd)
}
