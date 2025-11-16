package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_getPublicKeyCmd = &cobra.Command{
	Use:   "get-public-key",
	Short: "Gets information for the specified public key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_getPublicKeyCmd).Standalone()

	ivsRealtime_getPublicKeyCmd.Flags().String("arn", "", "ARN of the public key for which the information is to be retrieved.")
	ivsRealtime_getPublicKeyCmd.MarkFlagRequired("arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_getPublicKeyCmd)
}
