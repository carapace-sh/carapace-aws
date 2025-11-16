package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegatewayCmd = &cobra.Command{
	Use:   "storagegateway",
	Short: "Storage Gateway Service\n\nAmazon FSx File Gateway is no longer available to new customers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegatewayCmd).Standalone()

	rootCmd.AddCommand(storagegatewayCmd)
}
