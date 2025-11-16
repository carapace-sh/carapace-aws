package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_getIdNamespaceCmd = &cobra.Command{
	Use:   "get-id-namespace",
	Short: "Returns the `IdNamespace` with a given name, if it exists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_getIdNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_getIdNamespaceCmd).Standalone()

		entityresolution_getIdNamespaceCmd.Flags().String("id-namespace-name", "", "The name of the ID namespace.")
		entityresolution_getIdNamespaceCmd.MarkFlagRequired("id-namespace-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_getIdNamespaceCmd)
}
