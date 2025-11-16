package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_startCodegenJobCmd = &cobra.Command{
	Use:   "start-codegen-job",
	Short: "Starts a code generation job for a specified Amplify app and backend environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_startCodegenJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_startCodegenJobCmd).Standalone()

		amplifyuibuilder_startCodegenJobCmd.Flags().String("app-id", "", "The unique ID for the Amplify app.")
		amplifyuibuilder_startCodegenJobCmd.Flags().String("client-token", "", "The idempotency token used to ensure that the code generation job request completes only once.")
		amplifyuibuilder_startCodegenJobCmd.Flags().String("codegen-job-to-create", "", "The code generation job resource configuration.")
		amplifyuibuilder_startCodegenJobCmd.Flags().String("environment-name", "", "The name of the backend environment that is a part of the Amplify app.")
		amplifyuibuilder_startCodegenJobCmd.MarkFlagRequired("app-id")
		amplifyuibuilder_startCodegenJobCmd.MarkFlagRequired("codegen-job-to-create")
		amplifyuibuilder_startCodegenJobCmd.MarkFlagRequired("environment-name")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_startCodegenJobCmd)
}
