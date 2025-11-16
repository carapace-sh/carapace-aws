package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_configureAgentCmd = &cobra.Command{
	Use:   "configure-agent",
	Short: "Used by profiler agents to report their current state and to receive remote configuration updates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_configureAgentCmd).Standalone()

	codeguruprofiler_configureAgentCmd.Flags().String("fleet-instance-id", "", "A universally unique identifier (UUID) for a profiling instance.")
	codeguruprofiler_configureAgentCmd.Flags().String("metadata", "", "Metadata captured about the compute platform the agent is running on.")
	codeguruprofiler_configureAgentCmd.Flags().String("profiling-group-name", "", "The name of the profiling group for which the configured agent is collecting profiling data.")
	codeguruprofiler_configureAgentCmd.MarkFlagRequired("profiling-group-name")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_configureAgentCmd)
}
