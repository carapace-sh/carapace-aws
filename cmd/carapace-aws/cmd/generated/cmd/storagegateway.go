package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegatewayCmd = &cobra.Command{
	Use:   "storagegateway",
	Short: "Storage Gateway Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegatewayCmd).Standalone()

	})
	rootCmd.AddCommand(storagegatewayCmd)
}
