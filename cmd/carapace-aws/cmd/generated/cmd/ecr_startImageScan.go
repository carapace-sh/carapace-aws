package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_startImageScanCmd = &cobra.Command{
	Use:   "start-image-scan",
	Short: "Starts a basic image vulnerability scan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_startImageScanCmd).Standalone()

	ecr_startImageScanCmd.Flags().String("image-id", "", "")
	ecr_startImageScanCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry that contains the repository in which to start an image scan request.")
	ecr_startImageScanCmd.Flags().String("repository-name", "", "The name of the repository that contains the images to scan.")
	ecr_startImageScanCmd.MarkFlagRequired("image-id")
	ecr_startImageScanCmd.MarkFlagRequired("repository-name")
	ecrCmd.AddCommand(ecr_startImageScanCmd)
}
