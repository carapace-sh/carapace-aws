package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getKxConnectionStringCmd = &cobra.Command{
	Use:   "get-kx-connection-string",
	Short: "Retrieves a connection string for a user to connect to a kdb cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getKxConnectionStringCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_getKxConnectionStringCmd).Standalone()

		finspace_getKxConnectionStringCmd.Flags().String("cluster-name", "", "A name of the kdb cluster.")
		finspace_getKxConnectionStringCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_getKxConnectionStringCmd.Flags().String("user-arn", "", "The Amazon Resource Name (ARN) that identifies the user.")
		finspace_getKxConnectionStringCmd.MarkFlagRequired("cluster-name")
		finspace_getKxConnectionStringCmd.MarkFlagRequired("environment-id")
		finspace_getKxConnectionStringCmd.MarkFlagRequired("user-arn")
	})
	finspaceCmd.AddCommand(finspace_getKxConnectionStringCmd)
}
