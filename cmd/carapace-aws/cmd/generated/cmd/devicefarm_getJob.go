package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "Gets information about a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_getJobCmd).Standalone()

		devicefarm_getJobCmd.Flags().String("arn", "", "The job's ARN.")
		devicefarm_getJobCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_getJobCmd)
}
