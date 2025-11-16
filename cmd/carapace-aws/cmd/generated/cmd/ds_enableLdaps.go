package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_enableLdapsCmd = &cobra.Command{
	Use:   "enable-ldaps",
	Short: "Activates the switch for the specific directory to always use LDAP secure calls.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_enableLdapsCmd).Standalone()

	ds_enableLdapsCmd.Flags().String("directory-id", "", "The identifier of the directory.")
	ds_enableLdapsCmd.Flags().String("type", "", "The type of LDAP security to enable.")
	ds_enableLdapsCmd.MarkFlagRequired("directory-id")
	ds_enableLdapsCmd.MarkFlagRequired("type")
	dsCmd.AddCommand(ds_enableLdapsCmd)
}
