package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_deleteConstraintCmd = &cobra.Command{
	Use:   "delete-constraint",
	Short: "Deletes the specified constraint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_deleteConstraintCmd).Standalone()

	servicecatalog_deleteConstraintCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_deleteConstraintCmd.Flags().String("id", "", "The identifier of the constraint.")
	servicecatalog_deleteConstraintCmd.MarkFlagRequired("id")
	servicecatalogCmd.AddCommand(servicecatalog_deleteConstraintCmd)
}
