package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sts_getSessionTokenCmd = &cobra.Command{
	Use:   "get-session-token",
	Short: "Returns a set of temporary credentials for an Amazon Web Services account or IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sts_getSessionTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sts_getSessionTokenCmd).Standalone()

		sts_getSessionTokenCmd.Flags().String("duration-seconds", "", "The duration, in seconds, that the credentials should remain valid.")
		sts_getSessionTokenCmd.Flags().String("serial-number", "", "The identification number of the MFA device that is associated with the IAM user who is making the `GetSessionToken` call.")
		sts_getSessionTokenCmd.Flags().String("token-code", "", "The value provided by the MFA device, if MFA is required.")
	})
	stsCmd.AddCommand(sts_getSessionTokenCmd)
}
