package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_createWorkspaceCmd = &cobra.Command{
	Use:   "create-workspace",
	Short: "Creates a Prometheus workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_createWorkspaceCmd).Standalone()

	amp_createWorkspaceCmd.Flags().String("alias", "", "An alias that you assign to this workspace to help you identify it.")
	amp_createWorkspaceCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
	amp_createWorkspaceCmd.Flags().String("kms-key-arn", "", "(optional) The ARN for a customer managed KMS key to use for encrypting data within your workspace.")
	amp_createWorkspaceCmd.Flags().String("tags", "", "The list of tag keys and values to associate with the workspace.")
	ampCmd.AddCommand(amp_createWorkspaceCmd)
}
