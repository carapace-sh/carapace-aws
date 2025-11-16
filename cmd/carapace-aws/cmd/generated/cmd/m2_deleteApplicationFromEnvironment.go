package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_deleteApplicationFromEnvironmentCmd = &cobra.Command{
	Use:   "delete-application-from-environment",
	Short: "Deletes a specific application from the specific runtime environment where it was previously deployed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_deleteApplicationFromEnvironmentCmd).Standalone()

	m2_deleteApplicationFromEnvironmentCmd.Flags().String("application-id", "", "The unique identifier of the application you want to delete.")
	m2_deleteApplicationFromEnvironmentCmd.Flags().String("environment-id", "", "The unique identifier of the runtime environment where the application was previously deployed.")
	m2_deleteApplicationFromEnvironmentCmd.MarkFlagRequired("application-id")
	m2_deleteApplicationFromEnvironmentCmd.MarkFlagRequired("environment-id")
	m2Cmd.AddCommand(m2_deleteApplicationFromEnvironmentCmd)
}
