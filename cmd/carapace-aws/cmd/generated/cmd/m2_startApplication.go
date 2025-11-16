package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_startApplicationCmd = &cobra.Command{
	Use:   "start-application",
	Short: "Starts an application that is currently stopped.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_startApplicationCmd).Standalone()

	m2_startApplicationCmd.Flags().String("application-id", "", "The unique identifier of the application you want to start.")
	m2_startApplicationCmd.MarkFlagRequired("application-id")
	m2Cmd.AddCommand(m2_startApplicationCmd)
}
