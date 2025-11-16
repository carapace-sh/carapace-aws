package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getAccessKeyLastUsedCmd = &cobra.Command{
	Use:   "get-access-key-last-used",
	Short: "Retrieves information about when the specified access key was last used.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getAccessKeyLastUsedCmd).Standalone()

	iam_getAccessKeyLastUsedCmd.Flags().String("access-key-id", "", "The identifier of an access key.")
	iam_getAccessKeyLastUsedCmd.MarkFlagRequired("access-key-id")
	iamCmd.AddCommand(iam_getAccessKeyLastUsedCmd)
}
