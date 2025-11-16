package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getInstanceAccessDetailsCmd = &cobra.Command{
	Use:   "get-instance-access-details",
	Short: "Returns temporary SSH keys you can use to connect to a specific virtual private server, or *instance*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getInstanceAccessDetailsCmd).Standalone()

	lightsail_getInstanceAccessDetailsCmd.Flags().String("instance-name", "", "The name of the instance to access.")
	lightsail_getInstanceAccessDetailsCmd.Flags().String("protocol", "", "The protocol to use to connect to your instance.")
	lightsail_getInstanceAccessDetailsCmd.MarkFlagRequired("instance-name")
	lightsailCmd.AddCommand(lightsail_getInstanceAccessDetailsCmd)
}
