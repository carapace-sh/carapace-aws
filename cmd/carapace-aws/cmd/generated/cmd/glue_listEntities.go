package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listEntitiesCmd = &cobra.Command{
	Use:   "list-entities",
	Short: "Returns the available entities supported by the connection type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listEntitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listEntitiesCmd).Standalone()

		glue_listEntitiesCmd.Flags().String("catalog-id", "", "The catalog ID of the catalog that contains the connection.")
		glue_listEntitiesCmd.Flags().String("connection-name", "", "A name for the connection that has required credentials to query any connection type.")
		glue_listEntitiesCmd.Flags().String("data-store-api-version", "", "The API version of the SaaS connector.")
		glue_listEntitiesCmd.Flags().String("next-token", "", "A continuation token, included if this is a continuation call.")
		glue_listEntitiesCmd.Flags().String("parent-entity-name", "", "Name of the parent entity for which you want to list the children.")
	})
	glueCmd.AddCommand(glue_listEntitiesCmd)
}
