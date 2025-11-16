package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_putMetricPolicyCmd = &cobra.Command{
	Use:   "put-metric-policy",
	Short: "The metric policy that you want to add to the container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_putMetricPolicyCmd).Standalone()

	mediastore_putMetricPolicyCmd.Flags().String("container-name", "", "The name of the container that you want to add the metric policy to.")
	mediastore_putMetricPolicyCmd.Flags().String("metric-policy", "", "The metric policy that you want to associate with the container.")
	mediastore_putMetricPolicyCmd.MarkFlagRequired("container-name")
	mediastore_putMetricPolicyCmd.MarkFlagRequired("metric-policy")
	mediastoreCmd.AddCommand(mediastore_putMetricPolicyCmd)
}
