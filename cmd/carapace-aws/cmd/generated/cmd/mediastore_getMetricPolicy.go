package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_getMetricPolicyCmd = &cobra.Command{
	Use:   "get-metric-policy",
	Short: "Returns the metric policy for the specified container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_getMetricPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_getMetricPolicyCmd).Standalone()

		mediastore_getMetricPolicyCmd.Flags().String("container-name", "", "The name of the container that is associated with the metric policy.")
		mediastore_getMetricPolicyCmd.MarkFlagRequired("container-name")
	})
	mediastoreCmd.AddCommand(mediastore_getMetricPolicyCmd)
}
