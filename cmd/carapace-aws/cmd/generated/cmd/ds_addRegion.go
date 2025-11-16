package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_addRegionCmd = &cobra.Command{
	Use:   "add-region",
	Short: "Adds two domain controllers in the specified Region for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_addRegionCmd).Standalone()

	ds_addRegionCmd.Flags().String("directory-id", "", "The identifier of the directory to which you want to add Region replication.")
	ds_addRegionCmd.Flags().String("region-name", "", "The name of the Region where you want to add domain controllers for replication.")
	ds_addRegionCmd.Flags().String("vpcsettings", "", "")
	ds_addRegionCmd.MarkFlagRequired("directory-id")
	ds_addRegionCmd.MarkFlagRequired("region-name")
	ds_addRegionCmd.MarkFlagRequired("vpcsettings")
	dsCmd.AddCommand(ds_addRegionCmd)
}
