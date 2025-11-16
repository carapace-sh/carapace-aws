package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getRunCmd = &cobra.Command{
	Use:   "get-run",
	Short: "Gets information about a run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getRunCmd).Standalone()

	devicefarm_getRunCmd.Flags().String("arn", "", "The run's ARN.")
	devicefarm_getRunCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_getRunCmd)
}
