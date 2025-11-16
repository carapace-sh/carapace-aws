package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_deleteIdNamespaceCmd = &cobra.Command{
	Use:   "delete-id-namespace",
	Short: "Deletes the `IdNamespace` with a given name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_deleteIdNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_deleteIdNamespaceCmd).Standalone()

		entityresolution_deleteIdNamespaceCmd.Flags().String("id-namespace-name", "", "The name of the ID namespace.")
		entityresolution_deleteIdNamespaceCmd.MarkFlagRequired("id-namespace-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_deleteIdNamespaceCmd)
}
