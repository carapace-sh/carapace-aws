package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_batchCheckLayerAvailabilityCmd = &cobra.Command{
	Use:   "batch-check-layer-availability",
	Short: "Checks the availability of one or more image layers in a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_batchCheckLayerAvailabilityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_batchCheckLayerAvailabilityCmd).Standalone()

		ecr_batchCheckLayerAvailabilityCmd.Flags().String("layer-digests", "", "The digests of the image layers to check.")
		ecr_batchCheckLayerAvailabilityCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the image layers to check.")
		ecr_batchCheckLayerAvailabilityCmd.Flags().String("repository-name", "", "The name of the repository that is associated with the image layers to check.")
		ecr_batchCheckLayerAvailabilityCmd.MarkFlagRequired("layer-digests")
		ecr_batchCheckLayerAvailabilityCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_batchCheckLayerAvailabilityCmd)
}
