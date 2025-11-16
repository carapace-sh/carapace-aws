package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_setLocalConsolePasswordCmd = &cobra.Command{
	Use:   "set-local-console-password",
	Short: "Sets the password for your VM local console.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_setLocalConsolePasswordCmd).Standalone()

	storagegateway_setLocalConsolePasswordCmd.Flags().String("gateway-arn", "", "")
	storagegateway_setLocalConsolePasswordCmd.Flags().String("local-console-password", "", "The password you want to set for your VM local console.")
	storagegateway_setLocalConsolePasswordCmd.MarkFlagRequired("gateway-arn")
	storagegateway_setLocalConsolePasswordCmd.MarkFlagRequired("local-console-password")
	storagegatewayCmd.AddCommand(storagegateway_setLocalConsolePasswordCmd)
}
