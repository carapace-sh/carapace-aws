package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_getEnvironmentCmd = &cobra.Command{
	Use:   "get-environment",
	Short: "Describes a specific runtime environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_getEnvironmentCmd).Standalone()

	m2_getEnvironmentCmd.Flags().String("environment-id", "", "The unique identifier of the runtime environment.")
	m2_getEnvironmentCmd.MarkFlagRequired("environment-id")
	m2Cmd.AddCommand(m2_getEnvironmentCmd)
}
