package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_disableLdapsCmd = &cobra.Command{
	Use:   "disable-ldaps",
	Short: "Deactivates LDAP secure calls for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_disableLdapsCmd).Standalone()

	ds_disableLdapsCmd.Flags().String("directory-id", "", "The identifier of the directory.")
	ds_disableLdapsCmd.Flags().String("type", "", "The type of LDAP security to enable.")
	ds_disableLdapsCmd.MarkFlagRequired("directory-id")
	ds_disableLdapsCmd.MarkFlagRequired("type")
	dsCmd.AddCommand(ds_disableLdapsCmd)
}
