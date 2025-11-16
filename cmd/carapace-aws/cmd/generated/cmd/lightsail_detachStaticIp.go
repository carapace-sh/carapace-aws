package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_detachStaticIpCmd = &cobra.Command{
	Use:   "detach-static-ip",
	Short: "Detaches a static IP from the Amazon Lightsail instance to which it is attached.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_detachStaticIpCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_detachStaticIpCmd).Standalone()

		lightsail_detachStaticIpCmd.Flags().String("static-ip-name", "", "The name of the static IP to detach from the instance.")
		lightsail_detachStaticIpCmd.MarkFlagRequired("static-ip-name")
	})
	lightsailCmd.AddCommand(lightsail_detachStaticIpCmd)
}
