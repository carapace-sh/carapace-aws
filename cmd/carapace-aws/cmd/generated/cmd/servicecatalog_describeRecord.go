package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeRecordCmd = &cobra.Command{
	Use:   "describe-record",
	Short: "Gets information about the specified request operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeRecordCmd).Standalone()

	servicecatalog_describeRecordCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_describeRecordCmd.Flags().String("id", "", "The record identifier of the provisioned product.")
	servicecatalog_describeRecordCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
	servicecatalog_describeRecordCmd.Flags().String("page-token", "", "The page token for the next set of results.")
	servicecatalog_describeRecordCmd.MarkFlagRequired("id")
	servicecatalogCmd.AddCommand(servicecatalog_describeRecordCmd)
}
