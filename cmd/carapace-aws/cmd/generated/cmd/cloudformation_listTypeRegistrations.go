package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listTypeRegistrationsCmd = &cobra.Command{
	Use:   "list-type-registrations",
	Short: "Returns a list of registration tokens for the specified extension(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listTypeRegistrationsCmd).Standalone()

	cloudformation_listTypeRegistrationsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
	cloudformation_listTypeRegistrationsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	cloudformation_listTypeRegistrationsCmd.Flags().String("registration-status-filter", "", "The current status of the extension registration request.")
	cloudformation_listTypeRegistrationsCmd.Flags().String("type", "", "The kind of extension.")
	cloudformation_listTypeRegistrationsCmd.Flags().String("type-arn", "", "The Amazon Resource Name (ARN) of the extension.")
	cloudformation_listTypeRegistrationsCmd.Flags().String("type-name", "", "The name of the extension.")
	cloudformationCmd.AddCommand(cloudformation_listTypeRegistrationsCmd)
}
