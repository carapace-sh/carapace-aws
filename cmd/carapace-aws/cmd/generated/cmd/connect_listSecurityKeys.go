package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listSecurityKeysCmd = &cobra.Command{
	Use:   "list-security-keys",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listSecurityKeysCmd).Standalone()

	connect_listSecurityKeysCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listSecurityKeysCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listSecurityKeysCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listSecurityKeysCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listSecurityKeysCmd)
}
