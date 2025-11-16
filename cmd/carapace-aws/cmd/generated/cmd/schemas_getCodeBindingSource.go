package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_getCodeBindingSourceCmd = &cobra.Command{
	Use:   "get-code-binding-source",
	Short: "Get the code binding source URI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_getCodeBindingSourceCmd).Standalone()

	schemas_getCodeBindingSourceCmd.Flags().String("language", "", "The language of the code binding.")
	schemas_getCodeBindingSourceCmd.Flags().String("registry-name", "", "The name of the registry.")
	schemas_getCodeBindingSourceCmd.Flags().String("schema-name", "", "The name of the schema.")
	schemas_getCodeBindingSourceCmd.Flags().String("schema-version", "", "Specifying this limits the results to only this schema version.")
	schemas_getCodeBindingSourceCmd.MarkFlagRequired("language")
	schemas_getCodeBindingSourceCmd.MarkFlagRequired("registry-name")
	schemas_getCodeBindingSourceCmd.MarkFlagRequired("schema-name")
	schemasCmd.AddCommand(schemas_getCodeBindingSourceCmd)
}
