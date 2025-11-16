package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_startGuisessionCmd = &cobra.Command{
	Use:   "start-guisession",
	Short: "Initiates a graphical user interface (GUI) session that’s used to access a virtual computer’s operating system and application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_startGuisessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_startGuisessionCmd).Standalone()

		lightsail_startGuisessionCmd.Flags().String("resource-name", "", "The resource name.")
		lightsail_startGuisessionCmd.MarkFlagRequired("resource-name")
	})
	lightsailCmd.AddCommand(lightsail_startGuisessionCmd)
}
