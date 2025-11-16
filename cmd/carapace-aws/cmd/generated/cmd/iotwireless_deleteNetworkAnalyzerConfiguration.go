package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteNetworkAnalyzerConfigurationCmd = &cobra.Command{
	Use:   "delete-network-analyzer-configuration",
	Short: "Deletes a network analyzer configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteNetworkAnalyzerConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_deleteNetworkAnalyzerConfigurationCmd).Standalone()

		iotwireless_deleteNetworkAnalyzerConfigurationCmd.Flags().String("configuration-name", "", "")
		iotwireless_deleteNetworkAnalyzerConfigurationCmd.MarkFlagRequired("configuration-name")
	})
	iotwirelessCmd.AddCommand(iotwireless_deleteNetworkAnalyzerConfigurationCmd)
}
