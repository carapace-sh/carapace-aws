package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swf_registerDomainCmd = &cobra.Command{
	Use:   "register-domain",
	Short: "Registers a new domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swf_registerDomainCmd).Standalone()

	swf_registerDomainCmd.Flags().String("description", "", "A text description of the domain.")
	swf_registerDomainCmd.Flags().String("name", "", "Name of the domain to register.")
	swf_registerDomainCmd.Flags().String("tags", "", "Tags to be added when registering a domain.")
	swf_registerDomainCmd.Flags().String("workflow-execution-retention-period-in-days", "", "The duration (in days) that records and histories of workflow executions on the domain should be kept by the service.")
	swf_registerDomainCmd.MarkFlagRequired("name")
	swf_registerDomainCmd.MarkFlagRequired("workflow-execution-retention-period-in-days")
	swfCmd.AddCommand(swf_registerDomainCmd)
}
