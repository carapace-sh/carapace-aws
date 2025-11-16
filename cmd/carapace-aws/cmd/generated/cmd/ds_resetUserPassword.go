package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_resetUserPasswordCmd = &cobra.Command{
	Use:   "reset-user-password",
	Short: "Resets the password for any user in your Managed Microsoft AD or Simple AD directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_resetUserPasswordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_resetUserPasswordCmd).Standalone()

		ds_resetUserPasswordCmd.Flags().String("directory-id", "", "Identifier of the Managed Microsoft AD or Simple AD directory in which the user resides.")
		ds_resetUserPasswordCmd.Flags().String("new-password", "", "The new password that will be reset.")
		ds_resetUserPasswordCmd.Flags().String("user-name", "", "The user name of the user whose password will be reset.")
		ds_resetUserPasswordCmd.MarkFlagRequired("directory-id")
		ds_resetUserPasswordCmd.MarkFlagRequired("new-password")
		ds_resetUserPasswordCmd.MarkFlagRequired("user-name")
	})
	dsCmd.AddCommand(ds_resetUserPasswordCmd)
}
