package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_deletePublicKeyCmd = &cobra.Command{
	Use:   "delete-public-key",
	Short: "Deletes the specified public key used to sign stage participant tokens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_deletePublicKeyCmd).Standalone()

	ivsRealtime_deletePublicKeyCmd.Flags().String("arn", "", "ARN of the public key to be deleted.")
	ivsRealtime_deletePublicKeyCmd.MarkFlagRequired("arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_deletePublicKeyCmd)
}
