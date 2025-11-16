package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipmentCmd = &cobra.Command{
	Use:   "lookoutequipment",
	Short: "Amazon Lookout for Equipment is a machine learning service that uses advanced analytics to identify anomalies in machines from sensor data for use in predictive maintenance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipmentCmd).Standalone()

	rootCmd.AddCommand(lookoutequipmentCmd)
}
