package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_getApplicationVersionCmd = &cobra.Command{
	Use:   "get-application-version",
	Short: "Returns details about a specific version of a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_getApplicationVersionCmd).Standalone()

	m2_getApplicationVersionCmd.Flags().String("application-id", "", "The unique identifier of the application.")
	m2_getApplicationVersionCmd.Flags().String("application-version", "", "The specific version of the application.")
	m2_getApplicationVersionCmd.MarkFlagRequired("application-id")
	m2_getApplicationVersionCmd.MarkFlagRequired("application-version")
	m2Cmd.AddCommand(m2_getApplicationVersionCmd)
}
