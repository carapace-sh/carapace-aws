package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastoreDataCmd = &cobra.Command{
	Use:   "mediastore-data",
	Short: "An AWS Elemental MediaStore asset is an object, similar to an object in the Amazon S3 service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastoreDataCmd).Standalone()

	rootCmd.AddCommand(mediastoreDataCmd)
}
