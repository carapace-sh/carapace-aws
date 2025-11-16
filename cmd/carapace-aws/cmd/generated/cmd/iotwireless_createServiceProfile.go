package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createServiceProfileCmd = &cobra.Command{
	Use:   "create-service-profile",
	Short: "Creates a new service profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createServiceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_createServiceProfileCmd).Standalone()

		iotwireless_createServiceProfileCmd.Flags().String("client-request-token", "", "Each resource must have a unique client request token.")
		iotwireless_createServiceProfileCmd.Flags().String("lo-ra-wan", "", "The service profile information to use to create the service profile.")
		iotwireless_createServiceProfileCmd.Flags().String("name", "", "The name of the new resource.")
		iotwireless_createServiceProfileCmd.Flags().String("tags", "", "The tags to attach to the new service profile.")
	})
	iotwirelessCmd.AddCommand(iotwireless_createServiceProfileCmd)
}
