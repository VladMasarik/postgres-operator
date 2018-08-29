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

func ShowPolicy(httpclient *http.Client, APIServerURL, arg, BasicAuthUsername, BasicAuthPassword string) (msgs.ShowPolicyResponse, error) {

	var response msgs.ShowPolicyResponse

	url := APIServerURL + "/policies/" + arg + "?version=" + msgs.PGO_VERSION
	log.Debug("showPolicy called...[" + url + "]")

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
func CreatePolicy(httpclient *http.Client, APIServerURL, BasicAuthUsername, BasicAuthPassword string, request *msgs.CreatePolicyRequest) (msgs.CreatePolicyResponse, error) {

	var response msgs.CreatePolicyResponse

	jsonValue, _ := json.Marshal(request)
	url := APIServerURL + "/policies"
	log.Debug("createPolicy called...[" + url + "]")

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

func DeletePolicy(httpclient *http.Client, APIServerURL, arg, BasicAuthUsername, BasicAuthPassword string) (msgs.DeletePolicyResponse, error) {

	var response msgs.DeletePolicyResponse

	url := APIServerURL + "/policiesdelete/" + arg + "?version=" + msgs.PGO_VERSION

	log.Debug("delete policy called [" + url + "]")

	action := "GET"

	req, err := http.NewRequest(action, url, nil)
	if err != nil {
		return response, err
	}

	req.SetBasicAuth(BasicAuthUsername, BasicAuthPassword)

	resp, err := httpclient.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	log.Debugf("%v\n", resp)
	err = StatusCheck(resp)
	if err != nil {
		return response, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("%v\n", resp.Body)
		fmt.Println("Error: ", err)
		log.Println(err)
		return response, err
	}

	return response, err

}

func ApplyPolicy(httpclient *http.Client, APIServerURL, BasicAuthUsername, BasicAuthPassword string, request *msgs.ApplyPolicyRequest) (msgs.ApplyPolicyResponse, error) {

	var response msgs.ApplyPolicyResponse

	jsonValue, _ := json.Marshal(request)
	url := APIServerURL + "/policies/apply"
	log.Debug("applyPolicy called...[" + url + "]")

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
