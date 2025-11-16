package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwirelessCmd = &cobra.Command{
	Use:   "iotwireless",
	Short: "AWS IoT Wireless provides bi-directional communication between internet-connected wireless devices and the AWS Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwirelessCmd).Standalone()

	rootCmd.AddCommand(iotwirelessCmd)
}
