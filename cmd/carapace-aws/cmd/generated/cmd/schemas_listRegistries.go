package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_listRegistriesCmd = &cobra.Command{
	Use:   "list-registries",
	Short: "List the registries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_listRegistriesCmd).Standalone()

	schemas_listRegistriesCmd.Flags().String("limit", "", "")
	schemas_listRegistriesCmd.Flags().String("next-token", "", "The token that specifies the next page of results to return.")
	schemas_listRegistriesCmd.Flags().String("registry-name-prefix", "", "Specifying this limits the results to only those registry names that start with the specified prefix.")
	schemas_listRegistriesCmd.Flags().String("scope", "", "Can be set to Local or AWS to limit responses to your custom registries, or the ones provided by AWS.")
	schemasCmd.AddCommand(schemas_listRegistriesCmd)
}
