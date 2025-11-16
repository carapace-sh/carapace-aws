package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_deleteEnvironmentCmd = &cobra.Command{
	Use:   "delete-environment",
	Short: "Deletes a specific runtime environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_deleteEnvironmentCmd).Standalone()

	m2_deleteEnvironmentCmd.Flags().String("environment-id", "", "The unique identifier of the runtime environment you want to delete.")
	m2_deleteEnvironmentCmd.MarkFlagRequired("environment-id")
	m2Cmd.AddCommand(m2_deleteEnvironmentCmd)
}
