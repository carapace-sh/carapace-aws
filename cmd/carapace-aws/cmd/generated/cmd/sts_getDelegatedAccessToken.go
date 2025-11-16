package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_getDelegatedAccessTokenCmd = &cobra.Command{
	Use:   "get-delegated-access-token",
	Short: "This API is currently unavailable for general use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_getDelegatedAccessTokenCmd).Standalone()

	sts_getDelegatedAccessTokenCmd.Flags().String("trade-in-token", "", "")
	sts_getDelegatedAccessTokenCmd.MarkFlagRequired("trade-in-token")
	stsCmd.AddCommand(sts_getDelegatedAccessTokenCmd)
}
