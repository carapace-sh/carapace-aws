package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_getCodegenJobCmd = &cobra.Command{
	Use:   "get-codegen-job",
	Short: "Returns an existing code generation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_getCodegenJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_getCodegenJobCmd).Standalone()

		amplifyuibuilder_getCodegenJobCmd.Flags().String("app-id", "", "The unique ID of the Amplify app associated with the code generation job.")
		amplifyuibuilder_getCodegenJobCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app associated with the code generation job.")
		amplifyuibuilder_getCodegenJobCmd.Flags().String("id", "", "The unique ID of the code generation job.")
		amplifyuibuilder_getCodegenJobCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_getCodegenJobCmd.MarkFlagRequired("environment-name")
		amplifyuibuilder_getCodegenJobCmd.MarkFlagRequired("id")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_getCodegenJobCmd)
}
