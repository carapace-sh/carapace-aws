package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_enableSsoCmd = &cobra.Command{
	Use:   "enable-sso",
	Short: "Enables single sign-on for a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_enableSsoCmd).Standalone()

	ds_enableSsoCmd.Flags().String("directory-id", "", "The identifier of the directory for which to enable single-sign on.")
	ds_enableSsoCmd.Flags().String("password", "", "The password of an alternate account to use to enable single-sign on.")
	ds_enableSsoCmd.Flags().String("user-name", "", "The username of an alternate account to use to enable single-sign on.")
	ds_enableSsoCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_enableSsoCmd)
}
