package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_updatePrimaryRegionCmd = &cobra.Command{
	Use:   "update-primary-region",
	Short: "Changes the primary key of a multi-Region key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_updatePrimaryRegionCmd).Standalone()

	kms_updatePrimaryRegionCmd.Flags().String("key-id", "", "Identifies the current primary key.")
	kms_updatePrimaryRegionCmd.Flags().String("primary-region", "", "The Amazon Web Services Region of the new primary key.")
	kms_updatePrimaryRegionCmd.MarkFlagRequired("key-id")
	kms_updatePrimaryRegionCmd.MarkFlagRequired("primary-region")
	kmsCmd.AddCommand(kms_updatePrimaryRegionCmd)
}
