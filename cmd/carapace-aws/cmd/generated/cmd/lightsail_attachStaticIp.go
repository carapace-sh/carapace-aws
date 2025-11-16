package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_attachStaticIpCmd = &cobra.Command{
	Use:   "attach-static-ip",
	Short: "Attaches a static IP address to a specific Amazon Lightsail instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_attachStaticIpCmd).Standalone()

	lightsail_attachStaticIpCmd.Flags().String("instance-name", "", "The instance name to which you want to attach the static IP address.")
	lightsail_attachStaticIpCmd.Flags().String("static-ip-name", "", "The name of the static IP.")
	lightsail_attachStaticIpCmd.MarkFlagRequired("instance-name")
	lightsail_attachStaticIpCmd.MarkFlagRequired("static-ip-name")
	lightsailCmd.AddCommand(lightsail_attachStaticIpCmd)
}
