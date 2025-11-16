package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_enableClientAuthenticationCmd = &cobra.Command{
	Use:   "enable-client-authentication",
	Short: "Enables alternative client authentication methods for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_enableClientAuthenticationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_enableClientAuthenticationCmd).Standalone()

		ds_enableClientAuthenticationCmd.Flags().String("directory-id", "", "The identifier of the specified directory.")
		ds_enableClientAuthenticationCmd.Flags().String("type", "", "The type of client authentication to enable.")
		ds_enableClientAuthenticationCmd.MarkFlagRequired("directory-id")
		ds_enableClientAuthenticationCmd.MarkFlagRequired("type")
	})
	dsCmd.AddCommand(ds_enableClientAuthenticationCmd)
}
