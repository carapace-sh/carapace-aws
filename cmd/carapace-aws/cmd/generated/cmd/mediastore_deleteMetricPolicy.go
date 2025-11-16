package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_deleteMetricPolicyCmd = &cobra.Command{
	Use:   "delete-metric-policy",
	Short: "Deletes the metric policy that is associated with the specified container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_deleteMetricPolicyCmd).Standalone()

	mediastore_deleteMetricPolicyCmd.Flags().String("container-name", "", "The name of the container that is associated with the metric policy that you want to delete.")
	mediastore_deleteMetricPolicyCmd.MarkFlagRequired("container-name")
	mediastoreCmd.AddCommand(mediastore_deleteMetricPolicyCmd)
}
