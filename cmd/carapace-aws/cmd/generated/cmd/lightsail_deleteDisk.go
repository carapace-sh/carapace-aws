package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteDiskCmd = &cobra.Command{
	Use:   "delete-disk",
	Short: "Deletes the specified block storage disk.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteDiskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteDiskCmd).Standalone()

		lightsail_deleteDiskCmd.Flags().String("disk-name", "", "The unique name of the disk you want to delete (`my-disk`).")
		lightsail_deleteDiskCmd.Flags().String("force-delete-add-ons", "", "A Boolean value to indicate whether to delete all add-ons for the disk.")
		lightsail_deleteDiskCmd.MarkFlagRequired("disk-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteDiskCmd)
}
