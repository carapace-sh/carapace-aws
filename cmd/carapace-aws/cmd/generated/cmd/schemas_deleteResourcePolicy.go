package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Delete the resource-based policy attached to the specified registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_deleteResourcePolicyCmd).Standalone()

	schemas_deleteResourcePolicyCmd.Flags().String("registry-name", "", "The name of the registry.")
	schemasCmd.AddCommand(schemas_deleteResourcePolicyCmd)
}
