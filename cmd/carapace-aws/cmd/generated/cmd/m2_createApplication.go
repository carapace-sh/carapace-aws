package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates a new application with given parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_createApplicationCmd).Standalone()

		m2_createApplicationCmd.Flags().String("client-token", "", "A client token is a unique, case-sensitive string of up to 128 ASCII characters with ASCII values of 33-126 inclusive.")
		m2_createApplicationCmd.Flags().String("definition", "", "The application definition for this application.")
		m2_createApplicationCmd.Flags().String("description", "", "The description of the application.")
		m2_createApplicationCmd.Flags().String("engine-type", "", "The type of the target platform for this application.")
		m2_createApplicationCmd.Flags().String("kms-key-id", "", "The identifier of a customer managed key.")
		m2_createApplicationCmd.Flags().String("name", "", "The unique identifier of the application.")
		m2_createApplicationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) that identifies a role that the application uses to access Amazon Web Services resources that are not part of the application or are in a different Amazon Web Services account.")
		m2_createApplicationCmd.Flags().String("tags", "", "A list of tags to apply to the application.")
		m2_createApplicationCmd.MarkFlagRequired("definition")
		m2_createApplicationCmd.MarkFlagRequired("engine-type")
		m2_createApplicationCmd.MarkFlagRequired("name")
	})
	m2Cmd.AddCommand(m2_createApplicationCmd)
}
