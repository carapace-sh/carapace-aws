package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_enableAllFeaturesCmd = &cobra.Command{
	Use:   "enable-all-features",
	Short: "Enables all features in an organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_enableAllFeaturesCmd).Standalone()

	organizationsCmd.AddCommand(organizations_enableAllFeaturesCmd)
}
