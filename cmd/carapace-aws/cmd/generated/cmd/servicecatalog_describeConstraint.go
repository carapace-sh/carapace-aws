package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_describeConstraintCmd = &cobra.Command{
	Use:   "describe-constraint",
	Short: "Gets information about the specified constraint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_describeConstraintCmd).Standalone()

	servicecatalog_describeConstraintCmd.Flags().String("accept-language", "", "The language code.")
	servicecatalog_describeConstraintCmd.Flags().String("id", "", "The identifier of the constraint.")
	servicecatalog_describeConstraintCmd.MarkFlagRequired("id")
	servicecatalogCmd.AddCommand(servicecatalog_describeConstraintCmd)
}
