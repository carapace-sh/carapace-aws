package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createServiceCmd = &cobra.Command{
	Use:   "create-service",
	Short: "Create an Proton service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createServiceCmd).Standalone()

	proton_createServiceCmd.Flags().String("branch-name", "", "The name of the code repository branch that holds the code that's deployed in Proton.")
	proton_createServiceCmd.Flags().String("description", "", "A description of the Proton service.")
	proton_createServiceCmd.Flags().String("name", "", "The service name.")
	proton_createServiceCmd.Flags().String("repository-connection-arn", "", "The Amazon Resource Name (ARN) of the repository connection.")
	proton_createServiceCmd.Flags().String("repository-id", "", "The ID of the code repository.")
	proton_createServiceCmd.Flags().String("spec", "", "A link to a spec file that provides inputs as defined in the service template bundle schema file.")
	proton_createServiceCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton service.")
	proton_createServiceCmd.Flags().String("template-major-version", "", "The major version of the service template that was used to create the service.")
	proton_createServiceCmd.Flags().String("template-minor-version", "", "The minor version of the service template that was used to create the service.")
	proton_createServiceCmd.Flags().String("template-name", "", "The name of the service template that's used to create the service.")
	proton_createServiceCmd.MarkFlagRequired("name")
	proton_createServiceCmd.MarkFlagRequired("spec")
	proton_createServiceCmd.MarkFlagRequired("template-major-version")
	proton_createServiceCmd.MarkFlagRequired("template-name")
	protonCmd.AddCommand(proton_createServiceCmd)
}
