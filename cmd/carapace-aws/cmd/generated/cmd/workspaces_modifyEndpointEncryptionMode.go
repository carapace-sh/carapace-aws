package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifyEndpointEncryptionModeCmd = &cobra.Command{
	Use:   "modify-endpoint-encryption-mode",
	Short: "Modifies the endpoint encryption mode that allows you to configure the specified directory between Standard TLS and FIPS 140-2 validated mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifyEndpointEncryptionModeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspaces_modifyEndpointEncryptionModeCmd).Standalone()

		workspaces_modifyEndpointEncryptionModeCmd.Flags().String("directory-id", "", "The identifier of the directory.")
		workspaces_modifyEndpointEncryptionModeCmd.Flags().String("endpoint-encryption-mode", "", "The encryption mode used for endpoint connections when streaming to WorkSpaces Personal or WorkSpace Pools.")
		workspaces_modifyEndpointEncryptionModeCmd.MarkFlagRequired("directory-id")
		workspaces_modifyEndpointEncryptionModeCmd.MarkFlagRequired("endpoint-encryption-mode")
	})
	workspacesCmd.AddCommand(workspaces_modifyEndpointEncryptionModeCmd)
}
