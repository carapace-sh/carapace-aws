package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_describeCodeBindingCmd = &cobra.Command{
	Use:   "describe-code-binding",
	Short: "Describe the code binding URI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_describeCodeBindingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_describeCodeBindingCmd).Standalone()

		schemas_describeCodeBindingCmd.Flags().String("language", "", "The language of the code binding.")
		schemas_describeCodeBindingCmd.Flags().String("registry-name", "", "The name of the registry.")
		schemas_describeCodeBindingCmd.Flags().String("schema-name", "", "The name of the schema.")
		schemas_describeCodeBindingCmd.Flags().String("schema-version", "", "Specifying this limits the results to only this schema version.")
		schemas_describeCodeBindingCmd.MarkFlagRequired("language")
		schemas_describeCodeBindingCmd.MarkFlagRequired("registry-name")
		schemas_describeCodeBindingCmd.MarkFlagRequired("schema-name")
	})
	schemasCmd.AddCommand(schemas_describeCodeBindingCmd)
}
