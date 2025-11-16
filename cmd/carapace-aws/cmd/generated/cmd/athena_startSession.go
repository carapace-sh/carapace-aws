package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_startSessionCmd = &cobra.Command{
	Use:   "start-session",
	Short: "Creates a session for running calculations within a workgroup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_startSessionCmd).Standalone()

	athena_startSessionCmd.Flags().String("client-request-token", "", "A unique case-sensitive string used to ensure the request to create the session is idempotent (executes only once).")
	athena_startSessionCmd.Flags().String("description", "", "The session description.")
	athena_startSessionCmd.Flags().String("engine-configuration", "", "Contains engine data processing unit (DPU) configuration settings and parameter mappings.")
	athena_startSessionCmd.Flags().String("notebook-version", "", "The notebook version.")
	athena_startSessionCmd.Flags().String("session-idle-timeout-in-minutes", "", "The idle timeout in minutes for the session.")
	athena_startSessionCmd.Flags().String("work-group", "", "The workgroup to which the session belongs.")
	athena_startSessionCmd.MarkFlagRequired("engine-configuration")
	athena_startSessionCmd.MarkFlagRequired("work-group")
	athenaCmd.AddCommand(athena_startSessionCmd)
}
