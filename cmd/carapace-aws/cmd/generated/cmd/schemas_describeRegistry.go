package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_describeRegistryCmd = &cobra.Command{
	Use:   "describe-registry",
	Short: "Describes the registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_describeRegistryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_describeRegistryCmd).Standalone()

		schemas_describeRegistryCmd.Flags().String("registry-name", "", "The name of the registry.")
		schemas_describeRegistryCmd.MarkFlagRequired("registry-name")
	})
	schemasCmd.AddCommand(schemas_describeRegistryCmd)
}
