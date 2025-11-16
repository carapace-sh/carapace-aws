package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_postAgentProfileCmd = &cobra.Command{
	Use:   "post-agent-profile",
	Short: "Submits profiling data to an aggregated profile of a profiling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_postAgentProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_postAgentProfileCmd).Standalone()

		codeguruprofiler_postAgentProfileCmd.Flags().String("agent-profile", "", "The submitted profiling data.")
		codeguruprofiler_postAgentProfileCmd.Flags().String("content-type", "", "The format of the submitted profiling data.")
		codeguruprofiler_postAgentProfileCmd.Flags().String("profile-token", "", "Amazon CodeGuru Profiler uses this universally unique identifier (UUID) to prevent the accidental submission of duplicate profiling data if there are failures and retries.")
		codeguruprofiler_postAgentProfileCmd.Flags().String("profiling-group-name", "", "The name of the profiling group with the aggregated profile that receives the submitted profiling data.")
		codeguruprofiler_postAgentProfileCmd.MarkFlagRequired("agent-profile")
		codeguruprofiler_postAgentProfileCmd.MarkFlagRequired("content-type")
		codeguruprofiler_postAgentProfileCmd.MarkFlagRequired("profiling-group-name")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_postAgentProfileCmd)
}
