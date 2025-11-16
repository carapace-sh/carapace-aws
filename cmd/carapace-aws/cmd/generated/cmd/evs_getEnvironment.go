package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_getEnvironmentCmd = &cobra.Command{
	Use:   "get-environment",
	Short: "Returns a description of the specified environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_getEnvironmentCmd).Standalone()

	evs_getEnvironmentCmd.Flags().String("environment-id", "", "A unique ID for the environment.")
	evs_getEnvironmentCmd.MarkFlagRequired("environment-id")
	evsCmd.AddCommand(evs_getEnvironmentCmd)
}
