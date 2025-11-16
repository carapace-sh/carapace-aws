package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_disableClientAuthenticationCmd = &cobra.Command{
	Use:   "disable-client-authentication",
	Short: "Disables alternative client authentication methods for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_disableClientAuthenticationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_disableClientAuthenticationCmd).Standalone()

		ds_disableClientAuthenticationCmd.Flags().String("directory-id", "", "The identifier of the directory")
		ds_disableClientAuthenticationCmd.Flags().String("type", "", "The type of client authentication to disable.")
		ds_disableClientAuthenticationCmd.MarkFlagRequired("directory-id")
		ds_disableClientAuthenticationCmd.MarkFlagRequired("type")
	})
	dsCmd.AddCommand(ds_disableClientAuthenticationCmd)
}
