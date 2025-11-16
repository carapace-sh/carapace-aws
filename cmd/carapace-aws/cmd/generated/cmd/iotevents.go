package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ioteventsCmd = &cobra.Command{
	Use:   "iotevents",
	Short: "AWS IoT Events monitors your equipment or device fleets for failures or changes in operation, and triggers actions when such events occur.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ioteventsCmd).Standalone()

	rootCmd.AddCommand(ioteventsCmd)
}
