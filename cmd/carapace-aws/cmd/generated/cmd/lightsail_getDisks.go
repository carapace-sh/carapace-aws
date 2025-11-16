package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDisksCmd = &cobra.Command{
	Use:   "get-disks",
	Short: "Returns information about all block storage disks in your AWS account and region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDisksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getDisksCmd).Standalone()

		lightsail_getDisksCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getDisksCmd)
}
