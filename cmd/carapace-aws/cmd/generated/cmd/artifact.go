package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "This reference provides descriptions of the low-level AWS Artifact Service API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifactCmd).Standalone()

	rootCmd.AddCommand(artifactCmd)
}
