package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_startFailbackLaunchCmd = &cobra.Command{
	Use:   "start-failback-launch",
	Short: "Initiates a Job for launching the machine that is being failed back to from the specified Recovery Instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_startFailbackLaunchCmd).Standalone()

	drs_startFailbackLaunchCmd.Flags().String("recovery-instance-ids", "", "The IDs of the Recovery Instance whose failback launch we want to request.")
	drs_startFailbackLaunchCmd.Flags().String("tags", "", "The tags to be associated with the failback launch Job.")
	drs_startFailbackLaunchCmd.MarkFlagRequired("recovery-instance-ids")
	drsCmd.AddCommand(drs_startFailbackLaunchCmd)
}
