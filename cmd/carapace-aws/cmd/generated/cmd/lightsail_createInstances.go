package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createInstancesCmd = &cobra.Command{
	Use:   "create-instances",
	Short: "Creates one or more Amazon Lightsail instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createInstancesCmd).Standalone()

	lightsail_createInstancesCmd.Flags().String("add-ons", "", "An array of objects representing the add-ons to enable for the new instance.")
	lightsail_createInstancesCmd.Flags().String("availability-zone", "", "The Availability Zone in which to create your instance.")
	lightsail_createInstancesCmd.Flags().String("blueprint-id", "", "The ID for a virtual private server image (`app_wordpress_x_x` or `app_lamp_x_x`).")
	lightsail_createInstancesCmd.Flags().String("bundle-id", "", "The bundle of specification information for your virtual private server (or *instance*), including the pricing plan (`medium_x_x`).")
	lightsail_createInstancesCmd.Flags().String("custom-image-name", "", "(Discontinued) The name for your custom image.")
	lightsail_createInstancesCmd.Flags().String("instance-names", "", "The names to use for your new Lightsail instances.")
	lightsail_createInstancesCmd.Flags().String("ip-address-type", "", "The IP address type for the instance.")
	lightsail_createInstancesCmd.Flags().String("key-pair-name", "", "The name of your key pair.")
	lightsail_createInstancesCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
	lightsail_createInstancesCmd.Flags().String("user-data", "", "A launch script you can create that configures a server with additional user data.")
	lightsail_createInstancesCmd.MarkFlagRequired("availability-zone")
	lightsail_createInstancesCmd.MarkFlagRequired("blueprint-id")
	lightsail_createInstancesCmd.MarkFlagRequired("bundle-id")
	lightsail_createInstancesCmd.MarkFlagRequired("instance-names")
	lightsailCmd.AddCommand(lightsail_createInstancesCmd)
}
