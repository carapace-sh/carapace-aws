package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves the resource-based policy attached to a given registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_getResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_getResourcePolicyCmd).Standalone()

		schemas_getResourcePolicyCmd.Flags().String("registry-name", "", "The name of the registry.")
	})
	schemasCmd.AddCommand(schemas_getResourcePolicyCmd)
}
