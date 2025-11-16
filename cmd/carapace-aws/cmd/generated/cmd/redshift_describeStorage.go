package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeStorageCmd = &cobra.Command{
	Use:   "describe-storage",
	Short: "Returns account level backups storage size and provisional storage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeStorageCmd).Standalone()

	redshiftCmd.AddCommand(redshift_describeStorageCmd)
}
