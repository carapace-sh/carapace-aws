package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createGuisessionAccessDetailsCmd = &cobra.Command{
	Use:   "create-guisession-access-details",
	Short: "Creates two URLs that are used to access a virtual computer’s graphical user interface (GUI) session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createGuisessionAccessDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createGuisessionAccessDetailsCmd).Standalone()

		lightsail_createGuisessionAccessDetailsCmd.Flags().String("resource-name", "", "The resource name.")
		lightsail_createGuisessionAccessDetailsCmd.MarkFlagRequired("resource-name")
	})
	lightsailCmd.AddCommand(lightsail_createGuisessionAccessDetailsCmd)
}
