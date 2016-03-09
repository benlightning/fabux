// Copyright 2014 fabux authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package fabcore

import "flag"

func Flag() (*string, *bool, *string, *bool) {
	var (
		ConfigFile = flag.String("c", "server.ini", "General configuration file")
		canGlobal  = flag.Bool("g", true, "run global commands in the configuration file")
		HostLevel  = flag.String("h", "", "this parameter represents the server host in the configuration file,usage: -h=1,2,...")
		isLog      = flag.Bool("l", false, "This parameter defaults to indicate whether the print output execution results, usage: -l=true or -l")
	)
	flag.Parse()
	return ConfigFile, canGlobal, HostLevel, isLog
}
