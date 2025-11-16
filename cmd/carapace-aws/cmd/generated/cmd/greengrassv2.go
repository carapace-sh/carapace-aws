package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2Cmd = &cobra.Command{
	Use:   "greengrassv2",
	Short: "IoT Greengrass brings local compute, messaging, data management, sync, and ML inference capabilities to edge devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2Cmd).Standalone()

	rootCmd.AddCommand(greengrassv2Cmd)
}
