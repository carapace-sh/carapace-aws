package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listPublicKeysCmd = &cobra.Command{
	Use:   "list-public-keys",
	Short: "Gets summary information about all public keys in your account, in the AWS region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listPublicKeysCmd).Standalone()

	ivsRealtime_listPublicKeysCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	ivsRealtime_listPublicKeysCmd.Flags().String("next-token", "", "The first public key to retrieve.")
	ivsRealtimeCmd.AddCommand(ivsRealtime_listPublicKeysCmd)
}
