package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDiskCmd = &cobra.Command{
	Use:   "get-disk",
	Short: "Returns information about a specific block storage disk.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDiskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getDiskCmd).Standalone()

		lightsail_getDiskCmd.Flags().String("disk-name", "", "The name of the disk (`my-disk`).")
		lightsail_getDiskCmd.MarkFlagRequired("disk-name")
	})
	lightsailCmd.AddCommand(lightsail_getDiskCmd)
}
