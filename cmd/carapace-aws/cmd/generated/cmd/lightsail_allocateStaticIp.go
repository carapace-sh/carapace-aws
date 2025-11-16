package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_allocateStaticIpCmd = &cobra.Command{
	Use:   "allocate-static-ip",
	Short: "Allocates a static IP address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_allocateStaticIpCmd).Standalone()

	lightsail_allocateStaticIpCmd.Flags().String("static-ip-name", "", "The name of the static IP address.")
	lightsail_allocateStaticIpCmd.MarkFlagRequired("static-ip-name")
	lightsailCmd.AddCommand(lightsail_allocateStaticIpCmd)
}
