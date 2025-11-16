package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_releaseStaticIpCmd = &cobra.Command{
	Use:   "release-static-ip",
	Short: "Deletes a specific static IP from your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_releaseStaticIpCmd).Standalone()

	lightsail_releaseStaticIpCmd.Flags().String("static-ip-name", "", "The name of the static IP to delete.")
	lightsail_releaseStaticIpCmd.MarkFlagRequired("static-ip-name")
	lightsailCmd.AddCommand(lightsail_releaseStaticIpCmd)
}
