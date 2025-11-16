package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_updateConstraintCmd = &cobra.Command{
	Use:   "update-constraint",
	Short: "Updates the specified constraint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_updateConstraintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_updateConstraintCmd).Standalone()

		servicecatalog_updateConstraintCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_updateConstraintCmd.Flags().String("description", "", "The updated description of the constraint.")
		servicecatalog_updateConstraintCmd.Flags().String("id", "", "The identifier of the constraint.")
		servicecatalog_updateConstraintCmd.Flags().String("parameters", "", "The constraint parameters, in JSON format.")
		servicecatalog_updateConstraintCmd.MarkFlagRequired("id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_updateConstraintCmd)
}
