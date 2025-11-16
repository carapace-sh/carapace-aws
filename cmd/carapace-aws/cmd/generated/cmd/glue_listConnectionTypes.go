package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listConnectionTypesCmd = &cobra.Command{
	Use:   "list-connection-types",
	Short: "The `ListConnectionTypes` API provides a discovery mechanism to learn available connection types in Glue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listConnectionTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listConnectionTypesCmd).Standalone()

		glue_listConnectionTypesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		glue_listConnectionTypesCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	})
	glueCmd.AddCommand(glue_listConnectionTypesCmd)
}
