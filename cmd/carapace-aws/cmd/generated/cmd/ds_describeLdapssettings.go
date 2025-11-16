package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeLdapssettingsCmd = &cobra.Command{
	Use:   "describe-ldapssettings",
	Short: "Describes the status of LDAP security for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeLdapssettingsCmd).Standalone()

	ds_describeLdapssettingsCmd.Flags().String("directory-id", "", "The identifier of the directory.")
	ds_describeLdapssettingsCmd.Flags().String("limit", "", "Specifies the number of items that should be displayed on one page.")
	ds_describeLdapssettingsCmd.Flags().String("next-token", "", "The type of next token used for pagination.")
	ds_describeLdapssettingsCmd.Flags().String("type", "", "The type of LDAP security to enable.")
	ds_describeLdapssettingsCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_describeLdapssettingsCmd)
}
