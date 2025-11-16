package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_updateHybridAdCmd = &cobra.Command{
	Use:   "update-hybrid-ad",
	Short: "Updates the configuration of an existing hybrid directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_updateHybridAdCmd).Standalone()

	ds_updateHybridAdCmd.Flags().String("directory-id", "", "The identifier of the hybrid directory to update.")
	ds_updateHybridAdCmd.Flags().String("hybrid-administrator-account-update", "", "We create a hybrid directory administrator account when we create a hybrid directory.")
	ds_updateHybridAdCmd.Flags().String("self-managed-instances-settings", "", "Updates to the self-managed AD configuration, including DNS server IP addresses and Amazon Web Services System Manager managed node identifiers.")
	ds_updateHybridAdCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_updateHybridAdCmd)
}
