/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	_ "os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/BlackspaceInc/BlackspacePlatform/src/tools/service_generator/generator"
)

// generateBlackspaceServiceTemplateCmd represents the generateBlackspaceServiceTemplate command
var generateBlackspaceServiceTemplateCmd = &cobra.Command{
	Use:   "generateBlackspaceServiceTemplate",
	Short: "Generates a golang microservices",
	Long: `This tool is used to generate a microservice in golang specific to blackspace`,
	Example:`generateBlackspaceServiceTemplate --ApiType=REST --ServiceName=Shopper --ProtoDirectory=""
	         generateBlackspaceServiceTemplate --ApiType=GRAPHQL --ServiceName=Cart --ProtoDirectory="~/desktop/proto"`,
	RunE: generateServiceTemplate,
}


var (
	apiType string
	servicename string
	protocolBuffersDirectory string
)

const (
	REST string = "REST"
	GRAPHQL string = "GRAPHQL"
)

var generateServiceCmd = &cobra.Command{
	Use: "generateService",
	Short: "generates a golang microservice",
	Long: "generates a complete microservice with either a REST or GRAPHQL API interface",
}

func init() {

	generateBlackspaceServiceTemplateCmd.Flags().StringVarP(&apiType, "ApiType", "a","REST", "API interface")
	generateBlackspaceServiceTemplateCmd.Flags().StringVarP(&servicename, "ServiceName", "s", "", "Microservice Name")
	generateBlackspaceServiceTemplateCmd.Flags().StringVarP(&protocolBuffersDirectory, "ProtoDirectory", "d", "",
		"Directory at which protocol buffer files are located")
	rootCmd.AddCommand(generateBlackspaceServiceTemplateCmd)
}

func generateServiceTemplate(cmd *cobra.Command, args []string) error {
	if servicename == "" || apiType == ""{
		return fmt.Errorf("--ServiceName and --ApiType is required")
	}

	if apiType == GRAPHQL && protocolBuffersDirectory == "" {
		return fmt.Errorf("--ProtoDirectory must be present for --ApiType=GRAPHQL")
	}

	if strings.ToLower(apiType) == strings.ToLower(REST) {
		if err := generator.GenerateRESTMicroService(servicename, logger); err != nil {
			return err
		}
	}

	return nil
}

