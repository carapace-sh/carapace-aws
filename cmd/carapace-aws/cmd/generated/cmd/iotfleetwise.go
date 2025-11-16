package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwiseCmd = &cobra.Command{
	Use:   "iotfleetwise",
	Short: "Amazon Web Services IoT FleetWise is a fully managed service that you can use to collect, model, and transfer vehicle data to the Amazon Web Services cloud at scale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwiseCmd).Standalone()

	rootCmd.AddCommand(iotfleetwiseCmd)
}
