package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_getProgrammaticAccessCredentialsCmd = &cobra.Command{
	Use:   "get-programmatic-access-credentials",
	Short: "Request programmatic credentials to use with FinSpace SDK.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_getProgrammaticAccessCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_getProgrammaticAccessCredentialsCmd).Standalone()

		finspaceData_getProgrammaticAccessCredentialsCmd.Flags().String("duration-in-minutes", "", "The time duration in which the credentials remain valid.")
		finspaceData_getProgrammaticAccessCredentialsCmd.Flags().String("environment-id", "", "The FinSpace environment identifier.")
		finspaceData_getProgrammaticAccessCredentialsCmd.MarkFlagRequired("environment-id")
	})
	finspaceDataCmd.AddCommand(finspaceData_getProgrammaticAccessCredentialsCmd)
}
