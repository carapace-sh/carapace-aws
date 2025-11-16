package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_disableSsoCmd = &cobra.Command{
	Use:   "disable-sso",
	Short: "Disables single-sign on for a directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_disableSsoCmd).Standalone()

	ds_disableSsoCmd.Flags().String("directory-id", "", "The identifier of the directory for which to disable single-sign on.")
	ds_disableSsoCmd.Flags().String("password", "", "The password of an alternate account to use to disable single-sign on.")
	ds_disableSsoCmd.Flags().String("user-name", "", "The username of an alternate account to use to disable single-sign on.")
	ds_disableSsoCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_disableSsoCmd)
}
