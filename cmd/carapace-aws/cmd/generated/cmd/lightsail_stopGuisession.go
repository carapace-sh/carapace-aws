package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_stopGuisessionCmd = &cobra.Command{
	Use:   "stop-guisession",
	Short: "Terminates a web-based Amazon DCV session that’s used to access a virtual computer’s operating system or application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_stopGuisessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_stopGuisessionCmd).Standalone()

		lightsail_stopGuisessionCmd.Flags().String("resource-name", "", "The resource name.")
		lightsail_stopGuisessionCmd.MarkFlagRequired("resource-name")
	})
	lightsailCmd.AddCommand(lightsail_stopGuisessionCmd)
}
