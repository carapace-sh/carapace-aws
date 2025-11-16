package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getAssetCmd = &cobra.Command{
	Use:   "get-asset",
	Short: "Gets an Amazon DataZone asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getAssetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getAssetCmd).Standalone()

		datazone_getAssetCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain to which the asset belongs.")
		datazone_getAssetCmd.Flags().String("identifier", "", "The ID of the Amazon DataZone asset.")
		datazone_getAssetCmd.Flags().String("revision", "", "The revision of the Amazon DataZone asset.")
		datazone_getAssetCmd.MarkFlagRequired("domain-identifier")
		datazone_getAssetCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getAssetCmd)
}
