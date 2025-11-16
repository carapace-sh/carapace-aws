package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Describes the details of a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_getApplicationCmd).Standalone()

	m2_getApplicationCmd.Flags().String("application-id", "", "The identifier of the application.")
	m2_getApplicationCmd.MarkFlagRequired("application-id")
	m2Cmd.AddCommand(m2_getApplicationCmd)
}
