package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getInstanceCmd = &cobra.Command{
	Use:   "get-instance",
	Short: "Returns information about a specific Amazon Lightsail instance, which is a virtual private server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getInstanceCmd).Standalone()

		lightsail_getInstanceCmd.Flags().String("instance-name", "", "The name of the instance.")
		lightsail_getInstanceCmd.MarkFlagRequired("instance-name")
	})
	lightsailCmd.AddCommand(lightsail_getInstanceCmd)
}
