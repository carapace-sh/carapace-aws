package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_stopJobCmd = &cobra.Command{
	Use:   "stop-job",
	Short: "Initiates a stop request for the current job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_stopJobCmd).Standalone()

	devicefarm_stopJobCmd.Flags().String("arn", "", "Represents the Amazon Resource Name (ARN) of the Device Farm job to stop.")
	devicefarm_stopJobCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_stopJobCmd)
}
