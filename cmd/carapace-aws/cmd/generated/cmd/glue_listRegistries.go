package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listRegistriesCmd = &cobra.Command{
	Use:   "list-registries",
	Short: "Returns a list of registries that you have created, with minimal registry information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listRegistriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listRegistriesCmd).Standalone()

		glue_listRegistriesCmd.Flags().String("max-results", "", "Maximum number of results required per page.")
		glue_listRegistriesCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	})
	glueCmd.AddCommand(glue_listRegistriesCmd)
}
