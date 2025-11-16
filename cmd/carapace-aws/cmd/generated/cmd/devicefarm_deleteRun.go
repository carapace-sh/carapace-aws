package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_deleteRunCmd = &cobra.Command{
	Use:   "delete-run",
	Short: "Deletes the run, given the run ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_deleteRunCmd).Standalone()

	devicefarm_deleteRunCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) for the run to delete.")
	devicefarm_deleteRunCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_deleteRunCmd)
}
