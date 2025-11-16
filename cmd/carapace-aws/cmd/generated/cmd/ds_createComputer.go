package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_createComputerCmd = &cobra.Command{
	Use:   "create-computer",
	Short: "Creates an Active Directory computer object in the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_createComputerCmd).Standalone()

	ds_createComputerCmd.Flags().String("computer-attributes", "", "An array of [Attribute]() objects that contain any LDAP attributes to apply to the computer account.")
	ds_createComputerCmd.Flags().String("computer-name", "", "The name of the computer account.")
	ds_createComputerCmd.Flags().String("directory-id", "", "The identifier of the directory in which to create the computer account.")
	ds_createComputerCmd.Flags().String("organizational-unit-distinguished-name", "", "The fully-qualified distinguished name of the organizational unit to place the computer account in.")
	ds_createComputerCmd.Flags().String("password", "", "A one-time password that is used to join the computer to the directory.")
	ds_createComputerCmd.MarkFlagRequired("computer-name")
	ds_createComputerCmd.MarkFlagRequired("directory-id")
	ds_createComputerCmd.MarkFlagRequired("password")
	dsCmd.AddCommand(ds_createComputerCmd)
}
