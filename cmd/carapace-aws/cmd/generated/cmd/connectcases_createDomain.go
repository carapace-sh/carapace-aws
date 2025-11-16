package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_createDomainCmd = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates a domain, which is a container for all case data, such as cases, fields, templates and layouts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_createDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_createDomainCmd).Standalone()

		connectcases_createDomainCmd.Flags().String("name", "", "The name for your Cases domain.")
		connectcases_createDomainCmd.MarkFlagRequired("name")
	})
	connectcasesCmd.AddCommand(connectcases_createDomainCmd)
}
