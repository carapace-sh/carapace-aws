package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_cancelUpdateStackCmd = &cobra.Command{
	Use:   "cancel-update-stack",
	Short: "Cancels an update on the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_cancelUpdateStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_cancelUpdateStackCmd).Standalone()

		cloudformation_cancelUpdateStackCmd.Flags().String("client-request-token", "", "A unique identifier for this `CancelUpdateStack` request.")
		cloudformation_cancelUpdateStackCmd.Flags().String("stack-name", "", "If you don't pass a parameter to `StackName`, the API returns a response that describes all resources in the account.")
		cloudformation_cancelUpdateStackCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_cancelUpdateStackCmd)
}
