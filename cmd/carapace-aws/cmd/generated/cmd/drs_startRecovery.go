package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_startRecoveryCmd = &cobra.Command{
	Use:   "start-recovery",
	Short: "Launches Recovery Instances for the specified Source Servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_startRecoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_startRecoveryCmd).Standalone()

		drs_startRecoveryCmd.Flags().Bool("is-drill", false, "Whether this Source Server Recovery operation is a drill or not.")
		drs_startRecoveryCmd.Flags().Bool("no-is-drill", false, "Whether this Source Server Recovery operation is a drill or not.")
		drs_startRecoveryCmd.Flags().String("source-servers", "", "The Source Servers that we want to start a Recovery Job for.")
		drs_startRecoveryCmd.Flags().String("tags", "", "The tags to be associated with the Recovery Job.")
		drs_startRecoveryCmd.Flag("no-is-drill").Hidden = true
		drs_startRecoveryCmd.MarkFlagRequired("source-servers")
	})
	drsCmd.AddCommand(drs_startRecoveryCmd)
}
