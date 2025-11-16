package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeVolumesCmd = &cobra.Command{
	Use:   "describe-volumes",
	Short: "Describes one or more Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS volumes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeVolumesCmd).Standalone()

	fsx_describeVolumesCmd.Flags().String("filters", "", "Enter a filter `Name` and `Values` pair to view a select set of volumes.")
	fsx_describeVolumesCmd.Flags().String("max-results", "", "")
	fsx_describeVolumesCmd.Flags().String("next-token", "", "")
	fsx_describeVolumesCmd.Flags().String("volume-ids", "", "The IDs of the volumes whose descriptions you want to retrieve.")
	fsxCmd.AddCommand(fsx_describeVolumesCmd)
}
