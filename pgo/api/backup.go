package api

/*
 Copyright 2017-2018 Crunchy Data Solutions, Inc.
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

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	msgs "github.com/crunchydata/postgres-operator/apiservermsgs"
	"net/http"
)

func ShowBackup(httpclient *http.Client, APIServerURL, arg, BasicAuthUsername, BasicAuthPassword string) (msgs.ShowBackupResponse, error) {

	var response msgs.ShowBackupResponse

	url := APIServerURL + "/backups/" + arg + "?version=" + msgs.PGO_VERSION

	log.Debug("show backup called [" + url + "]")

	action := "GET"
	req, err := http.NewRequest(action, url, nil)
	if err != nil {
		return response, err
	}

	req.SetBasicAuth(BasicAuthUsername, BasicAuthPassword)
	resp, err := httpclient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error: Do: ", err)
		return response, err
	}
	log.Debugf("%v\n", resp)
	err = StatusCheck(resp)
	if err != nil {
		return response, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("%v\n", resp.Body)
		log.Println(err)
		return response, err
	}

	return response, err

}
func DeleteBackup(httpclient *http.Client, APIServerURL, arg, BasicAuthUsername, BasicAuthPassword string) (msgs.DeleteBackupResponse, error) {

	var response msgs.DeleteBackupResponse
	url := APIServerURL + "/backupsdelete/" + arg + "?version=" + msgs.PGO_VERSION

	log.Debug("delete backup called [" + url + "]")

	action := "GET"
	req, err := http.NewRequest(action, url, nil)
	if err != nil {
		return response, err
	}

	req.SetBasicAuth(BasicAuthUsername, BasicAuthPassword)
	resp, err := httpclient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error: Do: ", err)
		return response, err
	}
	log.Debugf("%v\n", resp)
	err = StatusCheck(resp)
	if err != nil {
		return response, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("%v\n", resp.Body)
		log.Println(err)
		return response, err
	}

	return response, err

}

func CreateBackup(httpclient *http.Client, APIServerURL, BasicAuthUsername, BasicAuthPassword string, request *msgs.CreateBackupRequest) (msgs.CreateBackupResponse, error) {

	var response msgs.CreateBackupResponse

	jsonValue, _ := json.Marshal(request)
	url := APIServerURL + "/backups"

	log.Debug("create backup called [" + url + "]")

	action := "POST"
	req, err := http.NewRequest(action, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return response, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(BasicAuthUsername, BasicAuthPassword)

	resp, err := httpclient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return response, err
	}

	log.Debugf("%v\n", resp)
	err = StatusCheck(resp)
	if err != nil {
		return response, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("%v\n", resp.Body)
		log.Println(err)
		return response, err
	}

	return response, err
}
