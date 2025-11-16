package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_getCallerIdentityCmd = &cobra.Command{
	Use:   "get-caller-identity",
	Short: "Returns details about the IAM user or role whose credentials are used to call the operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_getCallerIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sts_getCallerIdentityCmd).Standalone()

	})
	stsCmd.AddCommand(sts_getCallerIdentityCmd)
}
