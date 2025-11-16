package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_upgradeProfileVersionCmd = &cobra.Command{
	Use:   "upgrade-profile-version",
	Short: "Upgrade a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_upgradeProfileVersionCmd).Standalone()

	wellarchitected_upgradeProfileVersionCmd.Flags().String("client-request-token", "", "")
	wellarchitected_upgradeProfileVersionCmd.Flags().String("milestone-name", "", "")
	wellarchitected_upgradeProfileVersionCmd.Flags().String("profile-arn", "", "The profile ARN.")
	wellarchitected_upgradeProfileVersionCmd.Flags().String("workload-id", "", "")
	wellarchitected_upgradeProfileVersionCmd.MarkFlagRequired("profile-arn")
	wellarchitected_upgradeProfileVersionCmd.MarkFlagRequired("workload-id")
	wellarchitectedCmd.AddCommand(wellarchitected_upgradeProfileVersionCmd)
}
