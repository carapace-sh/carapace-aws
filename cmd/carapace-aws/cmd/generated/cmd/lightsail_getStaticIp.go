package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getStaticIpCmd = &cobra.Command{
	Use:   "get-static-ip",
	Short: "Returns information about an Amazon Lightsail static IP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getStaticIpCmd).Standalone()

	lightsail_getStaticIpCmd.Flags().String("static-ip-name", "", "The name of the static IP in Lightsail.")
	lightsail_getStaticIpCmd.MarkFlagRequired("static-ip-name")
	lightsailCmd.AddCommand(lightsail_getStaticIpCmd)
}
