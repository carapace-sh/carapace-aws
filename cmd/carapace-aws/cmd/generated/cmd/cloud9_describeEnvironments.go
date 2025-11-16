package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9_describeEnvironmentsCmd = &cobra.Command{
	Use:   "describe-environments",
	Short: "Gets information about Cloud9 development environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9_describeEnvironmentsCmd).Standalone()

	cloud9_describeEnvironmentsCmd.Flags().String("environment-ids", "", "The IDs of individual environments to get information about.")
	cloud9_describeEnvironmentsCmd.MarkFlagRequired("environment-ids")
	cloud9Cmd.AddCommand(cloud9_describeEnvironmentsCmd)
}
