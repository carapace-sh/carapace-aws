package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchangeCmd = &cobra.Command{
	Use:   "dataexchange",
	Short: "AWS Data Exchange is a service that makes it easy for AWS customers to exchange data in the cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchangeCmd).Standalone()

	rootCmd.AddCommand(dataexchangeCmd)
}
