// Copyright © 2017 Osiloke Emoekpere <me@osiloke.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"
	"strings"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/osiloke/mailservice/mailer"
	"github.com/spf13/cobra"
)

var redisURI string
var name string
var defaultQueue string

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// machineryCmd represents the machinery command
var machineryCmd = &cobra.Command{
	Use:   "machinery",
	Short: "Machinery service for sending emails",
	Long:  `Machinery service for sending emails`,
	Run: func(cmd *cobra.Command, args []string) {
		var cnf = &config.Config{
			Broker:        redisURI,
			DefaultQueue:  defaultQueue,
			ResultBackend: redisURI,
		}

		server, err := machinery.NewServer(cnf)
		if err != nil {
			// do something with the error
			if !strings.Contains(err.Error(), "interrupt") {
				panic(err)
			}
		}
		err = server.RegisterTask("mail", mailer.SendMail)
		if err != nil {
			panic(err)
		}
		worker := server.NewWorker(name, 10)
		err = worker.Launch()
		if err != nil {
			if !strings.Contains(err.Error(), "interrupt") {
				panic(err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(machineryCmd)
	machineryCmd.Flags().StringVarP(&name, "name", "n", getenv("NAME", "mailer"), "unique name of this worker")
	machineryCmd.Flags().StringVarP(&defaultQueue, "queue", "q", getenv("QUEUE", "machinery_tasks"), "queue for worker tasks")
	machineryCmd.Flags().StringVarP(&redisURI, "redis", "r", getenv("REDIS_URI", "redis://localhost:6379"), "redis uri")
}
