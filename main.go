package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var isEnablingSkip bool
	path := `.\BattleTech_Data\StreamingAssets\data\milestones\milestone_003_title_coronation_palace.json`

	milestone, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read file: %v", err)
		os.Exit(1)
	}

EnablingInput:
	for {
		fmt.Printf("\nTurn tutorial skip 'on' or 'off'?\n> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "on":
			isEnablingSkip = true
			break EnablingInput
		case "off":
			isEnablingSkip = false
			break EnablingInput
		default:
			continue EnablingInput
		}
	}

	milestoneStruct := Milestone003{}
	json.Unmarshal(milestone, &milestoneStruct)

	if isEnablingSkip {
		milestoneStruct.Results[0].Stats[0].Value = 114
	} else {
		milestoneStruct.Results[0].Stats[0].Value = 100
	}

	finished, err := json.MarshalIndent(milestoneStruct, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not prepare json file: %v", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(path, finished, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not save json file: %v", err)
		os.Exit(1)
	}

	fmt.Println("Done")
	fmt.Print("Press 'Enter' to close...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// Milestone003 is a struct containing our milestone json file's structure
type Milestone003 struct {
	Description struct {
		ID          string `json:"Id"`
		Name        string `json:"Name"`
		Details     string `json:"Details"`
		Icon        string `json:"Icon"`
		Cost        int    `json:"Cost"`
		Rarity      int    `json:"Rarity"`
		Purchasable bool   `json:"Purchasable"`
	} `json:"Description"`
	Scope        string `json:"Scope"`
	Requirements []struct {
		Scope           string `json:"Scope"`
		RequirementTags struct {
			Items            interface{} `json:"items"`
			TagSetSourceFile string      `json:"tagSetSourceFile"`
		} `json:"RequirementTags"`
		ExclusionTags struct {
			Items            interface{} `json:"items"`
			TagSetSourceFile string      `json:"tagSetSourceFile"`
		} `json:"ExclusionTags"`
		RequirementComparisons []struct {
			Obj           string      `json:"obj"`
			Op            string      `json:"op"`
			Val           int         `json:"val"`
			ValueConstant interface{} `json:"valueConstant"`
		} `json:"RequirementComparisons"`
	} `json:"Requirements"`
	Results []struct {
		Scope        string      `json:"Scope"`
		Requirements interface{} `json:"Requirements"`
		AddedTags    struct {
		} `json:"AddedTags"`
		RemovedTags struct {
		} `json:"RemovedTags"`
		Stats []struct {
			TypeString    string `json:"typeString"`
			Name          string `json:"name"`
			Value         int    `json:"value"`
			Set           string `json:"set"`
			ValueConstant string `json:"valueConstant"`
		} `json:"Stats"`
		Actions []struct {
			Type             string   `json:"Type"`
			Value            string   `json:"value"`
			AdditionalValues []string `json:"additionalValues"`
		} `json:"Actions"`
		ForceEvents     interface{} `json:"ForceEvents"`
		TemporaryResult bool        `json:"TemporaryResult"`
		ResultDuration  int         `json:"ResultDuration"`
	} `json:"Results"`
	Repeatable bool `json:"Repeatable"`
}
