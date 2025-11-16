package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_listIdNamespacesCmd = &cobra.Command{
	Use:   "list-id-namespaces",
	Short: "Returns a list of all ID namespaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_listIdNamespacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_listIdNamespacesCmd).Standalone()

		entityresolution_listIdNamespacesCmd.Flags().String("max-results", "", "The maximum number of `IdNamespace` objects returned per page.")
		entityresolution_listIdNamespacesCmd.Flags().String("next-token", "", "The pagination token from the previous API call.")
	})
	entityresolutionCmd.AddCommand(entityresolution_listIdNamespacesCmd)
}
