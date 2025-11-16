package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_putAccountPreferencesCmd = &cobra.Command{
	Use:   "put-account-preferences",
	Short: "Use this operation to set the account preference in the current Amazon Web Services Region to use long 17 character (63 bit) or short 8 character (32 bit) resource IDs for new EFS file system and mount target resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_putAccountPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_putAccountPreferencesCmd).Standalone()

		efs_putAccountPreferencesCmd.Flags().String("resource-id-type", "", "Specifies the EFS resource ID preference to set for the user's Amazon Web Services account, in the current Amazon Web Services Region, either `LONG_ID` (17 characters), or `SHORT_ID` (8 characters).")
		efs_putAccountPreferencesCmd.MarkFlagRequired("resource-id-type")
	})
	efsCmd.AddCommand(efs_putAccountPreferencesCmd)
}
