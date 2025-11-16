package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_startSessionCmd = &cobra.Command{
	Use:   "start-session",
	Short: "Initiates a connection to a target (for example, a managed node) for a Session Manager session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_startSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_startSessionCmd).Standalone()

		ssm_startSessionCmd.Flags().String("document-name", "", "The name of the SSM document you want to use to define the type of session, input parameters, or preferences for the session.")
		ssm_startSessionCmd.Flags().String("parameters", "", "The values you want to specify for the parameters defined in the Session document.")
		ssm_startSessionCmd.Flags().String("reason", "", "The reason for connecting to the instance.")
		ssm_startSessionCmd.Flags().String("target", "", "The managed node to connect to for the session.")
		ssm_startSessionCmd.MarkFlagRequired("target")
	})
	ssmCmd.AddCommand(ssm_startSessionCmd)
}
