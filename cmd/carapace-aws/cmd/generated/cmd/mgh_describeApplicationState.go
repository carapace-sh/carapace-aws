package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_describeApplicationStateCmd = &cobra.Command{
	Use:   "describe-application-state",
	Short: "Gets the migration status of an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_describeApplicationStateCmd).Standalone()

	mgh_describeApplicationStateCmd.Flags().String("application-id", "", "The configurationId in Application Discovery Service that uniquely identifies the grouped application.")
	mgh_describeApplicationStateCmd.MarkFlagRequired("application-id")
	mghCmd.AddCommand(mgh_describeApplicationStateCmd)
}
