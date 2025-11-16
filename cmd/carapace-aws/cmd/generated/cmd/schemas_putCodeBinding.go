package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_putCodeBindingCmd = &cobra.Command{
	Use:   "put-code-binding",
	Short: "Put code binding URI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_putCodeBindingCmd).Standalone()

	schemas_putCodeBindingCmd.Flags().String("language", "", "The language of the code binding.")
	schemas_putCodeBindingCmd.Flags().String("registry-name", "", "The name of the registry.")
	schemas_putCodeBindingCmd.Flags().String("schema-name", "", "The name of the schema.")
	schemas_putCodeBindingCmd.Flags().String("schema-version", "", "Specifying this limits the results to only this schema version.")
	schemas_putCodeBindingCmd.MarkFlagRequired("language")
	schemas_putCodeBindingCmd.MarkFlagRequired("registry-name")
	schemas_putCodeBindingCmd.MarkFlagRequired("schema-name")
	schemasCmd.AddCommand(schemas_putCodeBindingCmd)
}
