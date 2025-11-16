package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_getAccessKeyInfoCmd = &cobra.Command{
	Use:   "get-access-key-info",
	Short: "Returns the account identifier for the specified access key ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_getAccessKeyInfoCmd).Standalone()

	sts_getAccessKeyInfoCmd.Flags().String("access-key-id", "", "The identifier of an access key.")
	sts_getAccessKeyInfoCmd.MarkFlagRequired("access-key-id")
	stsCmd.AddCommand(sts_getAccessKeyInfoCmd)
}
