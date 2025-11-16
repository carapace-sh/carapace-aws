package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_describeEntityCmd = &cobra.Command{
	Use:   "describe-entity",
	Short: "Provides details regarding the entity used with the connection type, with a description of the data model for each field in the selected entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_describeEntityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_describeEntityCmd).Standalone()

		glue_describeEntityCmd.Flags().String("catalog-id", "", "The catalog ID of the catalog that contains the connection.")
		glue_describeEntityCmd.Flags().String("connection-name", "", "The name of the connection that contains the connection type credentials.")
		glue_describeEntityCmd.Flags().String("data-store-api-version", "", "The version of the API used for the data store.")
		glue_describeEntityCmd.Flags().String("entity-name", "", "The name of the entity that you want to describe from the connection type.")
		glue_describeEntityCmd.Flags().String("next-token", "", "A continuation token, included if this is a continuation call.")
		glue_describeEntityCmd.MarkFlagRequired("connection-name")
		glue_describeEntityCmd.MarkFlagRequired("entity-name")
	})
	glueCmd.AddCommand(glue_describeEntityCmd)
}
