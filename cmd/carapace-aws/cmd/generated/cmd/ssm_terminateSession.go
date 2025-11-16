package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_terminateSessionCmd = &cobra.Command{
	Use:   "terminate-session",
	Short: "Permanently ends a session and closes the data connection between the Session Manager client and SSM Agent on the managed node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_terminateSessionCmd).Standalone()

	ssm_terminateSessionCmd.Flags().String("session-id", "", "The ID of the session to terminate.")
	ssm_terminateSessionCmd.MarkFlagRequired("session-id")
	ssmCmd.AddCommand(ssm_terminateSessionCmd)
}
